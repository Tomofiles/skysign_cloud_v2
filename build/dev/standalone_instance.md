# シンプルでスタンドアロンなテスト/評価/開発用のローカル環境

_備考： すべてのデプロイ方法はDockerを用いています。_

## アーキテクチャ
アプリケーションの機能をテストしたい場合、Skysignはシンプルでスタンドアロンなサンドボックス環境を簡単に作成することができます。

サンドボックス環境は、データベースや負荷分散などを含んだ揮発性な環境であり、DockerおよびDocker Composeを用いて、いくつかのステップでコマンドを実行するだけで、試すことができます。

## クイックスタート
DockerおよびDocker Composeがインストール済みであれば、このフォルダ内で`docker-compose -f docker-compose_skysign-cloud-dev.yaml -p skysign-cloud-dev up`と実行するだけで、クラウド機能のサンドボックス環境を起動することができます。

なお、SkysignではUIの地図にCesiumを使用していますが、地形の表現に`Cesium World Terrain`というCesiumプロジェクトのサービスを使用しています。これは、Cesium ionというクラウドサービスの一部として提供されているもので、利用するにはアクセストークンを取得する必要があります。

Skysignでは、アクセストークンがなくてもサンドボックス環境は利用できるようになっていますが、地形表示は常時ONとなっていますので、ブラウザのコンソールログにエラーが出力されるようになります。（地形情報をCesium ionに取得しに行くが、認証エラーとなってしまうため）

もし、フル機能を確認したい場合は、アクセストークンを以下から取得してください。

[Terrain - cesium.com](https://cesium.com/docs/tutorials/terrain/)

取得したアクセストークンは、環境変数にエクスポートしてから、Docker Composeのビルドコマンドを実行してください。詳しくは、yamlファイルの冒頭コメントを参照ください。

## エッジについて
クラウド機能が構築できれば、エッジ機能からドローンをクラウドに接続できます。

ドローンとの接続は、Dronecodeプロジェクトの[Mavlink](https://github.com/mavlink/mavlink)のみが対応しており、[Mavsdk](https://github.com/mavlink/mavsdk)を用いたドローンとのアクセスが可能です。

エッジ機能のサンドボックス環境は、このフォルダ内で`docker-compose -f docker-compose_skysign-edge-dev.yaml -p skysign-edge-dev up`と実行するだけで起動することができます。その際に、クラウドの接続先を環境変数から指定してあげる必要があります。

Docker Composeの起動コマンドを実行する前に、各自の環境に合わせて`<ip>:<port>`の組み合わせで、`CLOUD_ADDRESS`というキーで環境変数を指定してください。詳しくは、yamlファイルの冒頭コメントを参照ください。

## OSSの使用について
エッジ機能のサンドボックス環境を構築するにあたり、以下のPX4ガゼボシミュレータのDockerイメージを使用しています。

[JonasVautherin/px4-gazebo-headless - github.com](https://github.com/JonasVautherin/px4-gazebo-headless)

PX4ガゼボシミュレータのホームポジションを環境変数で指定できるようになっています。`PX4_HOME_LAT`、`PX4_HOME_LON`、`PX4_HOME_ALT`を、それぞれ環境変数に設定してください。