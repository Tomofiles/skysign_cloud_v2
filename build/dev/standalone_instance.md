# シンプルでスタンドアロンなテスト/評価/開発用のローカルインスタンス

_備考： すべてのデプロイ方法はDockerを用いています。_

## アーキテクチャ

![Architecture diagram for running local processes](../../assets/generated/run_locally_architecture.png)

シンプルでスタンドアロンなアプリケーションの実行をテストしたい場合、Skysignは以下から構成される簡単なサンドボックス環境を作成することができます。
  * PostgreSQL: コンテナ上で実行されるPostgreSQLインスタンス。永続化のためのボリュームマウント等は考慮していないため、テスト/評価/開発以外の用途での使用は推奨されません。5432ポートでアクセスできます。
  * Static contents: Reactベースで構築されたUIを提供するWEBサーバ。3D地図ライブラリのCesiumを使用しており、Skysignのすべての機能を利用するには、Cesium ionのアクセストークンが必要です。5000ポートでアクセスできます。
  * Mission/Communication backend: Java（Spring Boot）で構築されたビジネスロジック・エンティティを提供するサーバ。5001ポートでgRPCアクセスができます。
  * gRPC gateway: go（grpc-gateway）で構築された中継サーバ。SkysignへのHTTPリクエストをバックエンドサーバのgRPCに翻訳するだけの、シンプルな役割となっています。ユーザー向けが8888ポート、ドローン向けが8889でアクセスできます。
  * Reverse proxy: nginxのリバースプロキシ機能を利用した、Pathベースの転送を提供するサーバ。8080ポートでアクセスできます。

## クイックスタート
DockerおよびDocker Composeがインストール済みであれば、このフォルダ内で`docker-compose -f docker-compose_skysign-dev.yaml -p skysign-dev up`と実行するだけで、サンドボックス環境を起動することができます。

なお、SkysignではUIの地図にCesiumを使用していますが、地形の表現に`Cesium World Terrain`というCesiumプロジェクトのサービスを使用しています。これは、Cesium ionというクラウドサービスの一部として提供されているもので、利用するにはアクセストークンを取得する必要があります。

Skysignでは、アクセストークンがなくてもサンドボックス環境は利用できるようになっていますが、地形表示は常時ONとなっていますので、ブラウザのコンソールログにエラーが出力されるようになります。（地形情報をCesium ionに取得しに行くが、認証エラーとなってしまうため）

もし、フル機能を確認したい場合は、アクセストークンを以下から取得してください。

[Terrain - cesium.com](https://cesium.com/docs/tutorials/terrain/)

取得したアクセストークンは、`REACT_APP_CESIUM_KEY`というキーで環境変数に設定してから、ビルドしてください。

## エッジについて
クラウド機能が構築できれば、エッジ機能からドローンをクラウドに接続できます。

現在、エッジ機能を簡単に試す環境を構築する手順を整理できていません。

ドローンとの接続は、Dronecodeプロジェクトの[Mavlink](https://github.com/mavlink/mavlink)のみが対応しており、[Mavsdk](https://github.com/mavlink/mavsdk)を用いたドローンとのアクセスが可能です。Dronecodeプロジェクト推奨のフライトコントローラ（[Pixhawk](https://pixhawk.org/)など）をお持ちの方は、Mavsdkでアクセスすることでお試しいただけるかと思います。
