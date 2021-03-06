# Skysign ドメイン・モデルの構築
## はじめに
Skysignの**Cloud**のBackendを構成する各種コンポーネントは、水平スケールに対応した、
マイクロサービス・アーキテクチャで構築されています。  
それぞれのマイクロサービスは、ドローンの運航にかかる重要な機能を有しており、
お互いに疎結合でありながら、お互いに協調しあって、サービス全体のコラボレーションを
実現しています。

当ドキュメントは、ドローン領域のビジネス・ドメインの整理・構築の過程で発見した、
いくつかのプラクティスを文書化し、公開することを目的に作成しています。  
当ドキュメントが、Skysignのアーキテクチャの理解を深める一助となることを、
期待しています。

## 目次
[【導入編】ドローン・ソフトウェアが解決するもの](./domain_models/01_introduction.md)

[【本編1】ドローンの機体との遠隔コミュニケーション](./domain_models/02_main_paper_1.md)

[【本編2】ウェイポイント飛行というドローンの重要任務](./domain_models/03_main_paper_2.md)

[【本編3】飛行の計画とドメイン間の参照関係](./domain_models/04_main_paper_3.md)

[【本編4】飛行の実行と依存リソースのカーボンコピー](./domain_models/05_main_paper_4.md)

[【本編5】飛行の実績と非同期的な飛行データの記録](./domain_models/06_main_paper_5.md)

[【まとめ編】境界付けられたコンテキストと統合](./domain_models/07_summary.md)
