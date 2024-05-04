# Gnuplotter
https://gnuplotapp2f.trap.show/

## 画面
<p align="center">
  <img src="https://github.com/Luftalian/gnuplot_app2/assets/105796502/e8c137d7-75d1-498e-94b0-35ca63598d48" width=70% position=center>
</p>

## 開発目的
gnuplotのグラフをWeb上で作りたい。iPadとかで簡単に編集したい。かつ、軸ラベルの有無などの細かい設定も変更できるようにしたい。

## 開発方法
```bash
git clone https://github.com/Luftalian/gnuplot_app2.git
cd gnuplot_app2
docker compose up --build
```

バックエンドが http://localhost:8080 に

dbが http://localhost:8081 に

フロントエンドが http://localhost:8082 に
起動します。

## 技術スタック
- Docker
- Go (backend)
- Vue.js (frontend)
- MySQL (db)
- Neoshowcase (deploy)

## 今後
- データファイルアップロード
- User機能
- スクリプトを直接変更
- Chatサポート
