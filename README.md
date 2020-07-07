# Skysign Cloud (version 2)

Skysignは、ウェブブラウザから簡潔なUIで、ドローンの飛行制御・監視が行えるプラットフォームです。  
このリポジトリは、[Skysign Cloud β版](https://github.com/Tomofiles/skysign_cloud)と同等の機能を、
新しいアーキテクチャで刷新したものです。

version 2の狙いは、β版をスケールしていく際に発覚した、いくつかの課題を解消することです。
- websocket、server-sent eventsと、コンテナ系マネージドサービスとの相性の悪さ
- websocket、server-sent eventsによる、スケールアウト実現の難化
- フロントエンドのソースコードのカオス化

version 2は全面的に、InterUSSプロジェクトの[DSS](https://github.com/interuss/dss)に影響を受けて、
アーキテクチャやコンセプトが構築されています。  
DSSは、欧米における無人航空機（UAS）の社会実装を、テクノロジー面から実現するために活動している、
OSSコミュニティです。  
（なお、当プロジェクトおよびTomofilesは、InterUSSと一切の関係はありません。）

## デモ
![DEMO](https://user-images.githubusercontent.com/27773127/86793686-aa26e000-c0a6-11ea-9fa1-ef1fdf6706a6.gif)

## コンセプト
Skysignが提供する各種コンポーネントと、コンポーネント間のコミュニケーションの概要については、
[こちら](concepts.md)をご覧ください。

## アーキテクチャ概要
### Overview
![Simplified architecture diagram](assets/generated/simple_architecture.png)

Skysignは、大きく分けて`Cloud`と`Edge`の2つのコンポーネントから構成されます。

`Cloud`は、クラウドプラットフォーム上にデプロイされる、コンテナライズされた分散アプリケーションです。
すべてのコンテナコンポーネントは、HTTPおよびgRPCの通信プロトコルを採用しており、
負荷分散を目的とした水平スケールに対応しています。
また、Reactベースで構築されたUIをユーザーに提供し、RestfulなAPIを通してサーバーと情報を交換します。

`Edge`は、ドローンの機体に搭載されるコンパニオン・コンピュータにデプロイされ、
モバイル通信網等を介してサーバーと通信を行う、`Internet of Drones`の実装です。

ユーザー、ドローン、そしてクラウドのそれぞれのコンポーネント間は、
リソース指向なAPIにより疎結合な接続を維持し、かつ
高度な自動化を実現し、インターネットを介したドローンの遠隔的な制御・監視をサポートします。

### Reverse proxy
ユーザーおよびドローンからのアクセスのエントリーポイントで、
分散するバッグエンドサーバーとの中継や、負荷分散を担います。

### gRPC gateway
ユーザーおよびドローンからのHTTPリクエストをシンプルに翻訳し、
gRPCにて提供されるビジネスロジックと中継するサーバーコンポーネント。
当コンポーネントのソースコードは、[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)によりprotoファイルから自動生成されます。

### Static contents
[React](https://ja.reactjs.org/)をベース、[Material-UI](https://material-ui.com/)をビュー、
[Cesium](https://cesium.com/cesiumjs/)を地図エンジンに利用し、直感的なUIを構築して提供します。

### Mission/Communication backend
Skysignが提供するビジネスロジックおよびエンティティを提供するコンポーネント。
[Spring Boot](https://spring.io/projects/spring-boot)をベースに、
DDD(ドメイン駆動設計)を採用したマイクロサービス・アーキテクチャを目指しています。

### PostgreSQL (DB)
ビジネスロジックを支えるエンティティを格納するデータベース。

### Edge
ドローンのコンパニオン・コンピュータ（Raspberry Pi等）上で実行する常駐アプリケーション。
ドローンのフライトコントローラとしてDronecodeの[PX4](https://px4.io/)に対応しており、
[Mavsdk](https://github.com/mavlink/MAVSDK)を介してPX4と通信を行い、
ドローンのコントロールとテレメトリーの収集、及びクラウドとの通信を行います。

## ライセンス
MIT License

## 免責
当リポジトリーの利用は自由ですが、現在、絶賛開発中につき、多くの不具合を抱えています。  
当リポジトリーを使用したドローンの飛行にかかる、故障や損傷、及び本人や第三者への損害に関しては、
一切の責任を負いかねますので、あらかじめご了承ください。