---
name: test-create-vitest
description: プロジェクトのvitest単体テストを作成するスキル。ユーザーが「テストを書いて」「単体テストを作成して」「vitestでテストして」などと言った場合、または特定のモジュールやコンポーネントのテストが必要な場合に使用する。
---

# Vitest単体テスト作成スキル

## このスキルの使い方

ユーザーがテスト作成を依頼したら、以下の手順で進める：

1. **テスト対象の特定**: ユーザーが指定したファイル、関数、またはモジュールを確認する
2. **コードの分析**: テスト対象のソースコードを読み込み、エクスポート、依存関係、分岐条件を把握する
3. **テストケースの設計**: コードの実装に基づいて、正常系・異常系・境界値のテストケースを設計する
4. **テストファイルの作成**: 以下の規約に従ってテストを実装する

## テストファイル配置規約

TBD

## テスト設計の基本原則

### Arrange-Act-Assertパターン

```typescript
it("テストケース名", async () => {
  // Arrange: テストデータとモックをセットアップ
  vi.mocked(dependency).mockResolvedValue(testData);

  // Act: テスト対象の関数を実行
  const result = await targetFunction(params);

  // Assert: 期待値と実際の結果を検証
  expect(result).toEqual(expectedValue);
});
```

### テストの命名規則

- **describeブロック**: 関数名 → 機能カテゴリの順にネスト
- **itブロック**: 何をテストするのか明確に記述（日本語可）
- **コメント**: `// 正常系`、`// 異常系`、`// 境界値` で種別を明示

### モックとクリーンアップ

```typescript
import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";

vi.mock("~/models/[modelName].server");

describe("テストスイート", () => {
  afterEach(() => {
    vi.clearAllMocks();
  });
});
```

### テストデータの設計原則

重複定義を避け、保守性を高めるためのルール

#### ベースデータパターンを使うべきケース

以下の条件を**すべて満たす**場合、ベースデータパターンを使用する：

1. **複数のテストケースで同じ構造のデータを使う**
2. **各テストケースで一部のフィールドだけを変更する**
3. **データ構造が3フィールド以上ある**

#### ベースデータの配置ルール（スコープに応じた定義位置）

```typescript
// ✅ Good: テストファイル全体で共有する場合は最上位に定義
const baseData = {
  id: "test-id",
  name: "test-name",
  createdAt: new Date("2024-01-01"),
};

describe("functionName", () => {
  it("正常系テスト", () => {
    const testData = { ...baseData };
    // ...
  });
});

describe("anotherFunction", () => {
  it("別の関数でも同じデータを使用", () => {
    const testData = { ...baseData, name: "updated" };
    // ...
  });
});
```

```typescript
// ✅ Good: 特定のdescribeブロック内でのみ共有する場合はdescribe直下に定義
describe("functionName", () => {
  const baseData = {
    id: "test-id",
    name: "test-name",
    createdAt: new Date("2024-01-01"),
  };

  it("正常系テスト", () => {
    const testData = { ...baseData };
    // ...
  });

  it("異常系テスト", () => {
    const testData = { ...baseData, name: null };
    // ...
  });
});

describe("anotherFunction", () => {
  // このdescribeブロックでは上記のbaseDataは参照できない（スコープ外）
  const differentBaseData = {
    email: "test@example.com",
    age: 20,
  };
  // ...
});
```

#### ベースデータパターンを使わないケース

以下の場合は、各テストケース内でデータを直接定義する：

```typescript
// ❌ Bad: 1〜2個のフィールドしかない場合は不要
const baseData = { id: "test-id" };

// ✅ Good: シンプルなデータは直接定義
it("テスト", () => {
  const testData = { id: "test-id" };
  // ...
});
```

```typescript
// ❌ Bad: 各テストで完全に異なるデータを使う場合は不要
const baseData = { name: "test" };

it("テスト1", () => {
  const testData = { email: "test@example.com", age: 20 }; // baseData未使用
  // ...
});
```

#### 期待される効果

- **行数削減**: 同じデータ構造を繰り返し書かない
- **保守性向上**: データ構造変更時の修正箇所が1箇所になる
- **可読性向上**: 各テストで「何が違うのか」が明確になる

## テストパターンリファレンス

### バリデーションのテストパターン

```typescript
import { describe, it, expect } from "vitest";
import { validator } from "~/components/validators/[ValidatorName]";

describe("validator", () => {
  // 正常系
  it("有効な入力でエラーが発生しないこと", async () => {
    const formData = new FormData();
    formData.set("fieldName", "valid-value");
    const result = await validator({}).validate(formData);
    expect(result.error).toBeUndefined();
  });

  // 異常系
  it("空文字列でエラーが発生すること", async () => {
    const formData = new FormData();
    formData.set("fieldName", "");
    const result = await validator({}).validate(formData);
    expect(result.error?.fieldErrors.fieldName).toBeDefined();
  });

  // 境界値
  it("最大文字数超過でエラーが発生すること", async () => {
    const formData = new FormData();
    formData.set("fieldName", "a".repeat(101));
    const result = await validator({}).validate(formData);
    expect(result.error?.fieldErrors.fieldName).toContain("at most");
  });
});
```

### サーバーサイドロジックのテストパターン

```typescript
import { describe, it, expect, vi, afterEach } from "vitest";
import { targetFunction } from "~/services/[serviceName].server";
import * as dependencyModel from "~/models/[modelName].server";

vi.mock("~/models/[modelName].server");

// ベースデータ（各テストでスプレッド構文で継承）
const baseData = {
  id: "test-id",
  name: "test-name",
  createdAt: new Date("2024-01-01"),
};

afterEach(() => vi.clearAllMocks());

describe("[functionName]", () => {
  // 正常系
  it("データを正しく取得すること", async () => {
    vi.mocked(dependencyModel.getData).mockResolvedValue([baseData]);
    const result = await targetFunction({ id: "test-id" });
    expect(result).toHaveLength(1);
    expect(result[0].name).toBe("test-name");
  });

  // 異常系
  it("空データに対して空配列を返すこと", async () => {
    vi.mocked(dependencyModel.getData).mockResolvedValue([]);
    const result = await targetFunction({ id: "test-id" });
    expect(result).toEqual([]);
  });

  // 境界値
  it("nullフィールドを正しく処理すること", async () => {
    vi.mocked(dependencyModel.getData).mockResolvedValue([
      { ...baseData, name: null as any }
    ]);
    const result = await targetFunction({ id: "test-id" });
    expect(result[0].name).toBeNull();
  });
});
```

### Reactコンポーネントのテストパターン

```typescript
import "@testing-library/jest-dom";
import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { ComponentName } from "~/components/[ComponentName]";

describe("ComponentName", () => {
  const defaultProps = {
    label: "Label",
    onSubmit: vi.fn(),
  };

  beforeEach(() => vi.clearAllMocks());

  // 正常系: レンダリング
  it("propsに基づいて正しくレンダリングされること", () => {
    render(<ComponentName {...defaultProps} />);
    expect(screen.getByText("Label")).toBeInTheDocument();
  });

  // 正常系: インタラクション
  it("ボタンクリックでonSubmitが呼ばれること", () => {
    render(<ComponentName {...defaultProps} />);
    fireEvent.click(screen.getByRole("button"));
    expect(defaultProps.onSubmit).toHaveBeenCalledTimes(1);
  });

  // 異常系: エラー表示
  it("エラーがある場合メッセージを表示すること", () => {
    render(<ComponentName {...defaultProps} error="Error occurred" />);
    expect(screen.getByText("Error occurred")).toBeInTheDocument();
  });
});
```

## テストケース設計の考え方

### コード分析時のチェックポイント

1. **関数の引数**: 必須/オプション、型、デフォルト値
2. **戻り値**: 型、取りうる値の範囲
3. **分岐条件**: if文、switch文、三項演算子
4. **外部依存**: モックが必要なモジュール

### テストケースの分類

- **正常系**: 期待される入力に対して正しい出力を返すこと
- **異常系**: 無効な入力、エッジケース（null、undefined、空文字列など）
- **境界値**: 最小値・最大値、配列の0件・1件・複数件

## テスト実行コマンド

```bash
yarn test --run                                    # 全テスト実行
yarn test --run vitest/app/services/foo.test.ts   # 特定ファイル
yarn test --run -t "functionName"                  # 特定のdescribe
yarn test --coverage                               # カバレッジ
```
