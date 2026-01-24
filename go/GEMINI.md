日本語で対話すること

# プロジェクト概要

このリポジトリは、Goバックエンド学習用のプレイグラウンドです。Goでバックエンドサービスを構築するためのさまざまな概念やライブラリを実演する、いくつかの小さく独立したGoプロジェクトが含まれています。

## プロジェクト

### 01_simple

*   **目的:** `net/http`パッケージを使用したシンプルなAPIサーバーの構築方法を示します。
*   **主要な概念:**
    *   基本的なリクエスト処理
    *   `http.ResponseWriter`と`http.Request`
    *   ルーティングのための`http.ServeMux`
    *   デフォルトとカスタムのマルチプレクサ
*   **プロジェクトの実行方法:**
    ```bash
    go run main.go
    ```
*   **テスト:**
    ```bash
    curl http://localhost:8080/api/hello
    ```

### 02_docker

*   **目的:** GoアプリケーションをDockerを使用してコンテナ化する方法を示します。
*   **主要な概念:**
    *   Goプロジェクト用の`Dockerfile`の記述
    *   Dockerコンテナのビルドと実行
    *   GoアプリケーションをDocker化するためのベストプラクティス（例: `CGO_ENABLED=0`、`ENTRYPOINT` vs `CMD`）
*   **プロジェクトのビルドと実行:**
    ```bash
    docker build . -t go-hello
    docker run -d -p 8080:8080 go-hello
    ```
*   **テスト:**
    ```bash
    curl -i http://localhost:8080/api/hello
    ```

### 03_request-handling

*   **目的:** Go APIサーバーで受信リクエストを処理するさまざまな方法を示します。
*   **主要な概念:**
    *   クエリパラメータの処理
    *   URLパスパラメータの処理
    *   JSONリクエストボディの処理
    *   基本的なエラー処理
*   **プロジェクトの実行方法:**
    ```bash
    go run main.go
    ```
*   **テスト:**
    *   **クエリパラメータ:** `curl "http://localhost:8080/api/greet?name=John"`
    *   **パスパラメータ:** `curl http://localhost:8080/api/user/123`
    *   **リクエストボディ:** `curl -X POST -d '{"message":"hello"}' http://localhost:8080/api/echo`

### 10_go-mid

*   **目的:** `gorilla/mux`ルーターを使用して構築されたシンプルなブログのようなAPI。
*   **主要な概念:**
    *   サードパーティ製ルーター（`gorilla/mux`）の使用
    *   シンプルなRESTfulサービスのためのAPIルートの定義
    *   異なるHTTPメソッド（GET、POST）の処理
*   **プロジェクトの実行方法:**
    ```bash
    go run main.go
    ```
*   **テスト:**
    *   `curl http://localhost:8080/hello`
    *   `curl http://localhost:8080/article/list`
    *   `curl http://localhost:8080/article/123`
    *   `curl -X POST http://localhost:8080/article`

## 開発規約

*   各サブディレクトリは、独自の`go.mod`ファイルを持つ自己完結型のGoプロジェクトです。
*   プロジェクトは一般的に標準的なGoコーディング規約に従っています。
*   `mise.toml`の使用は、Goのバージョン管理を含むプロジェクト固有の開発環境の管理に[mise](https://mise.jdx.dev/)が使用されていることを示唆しています。