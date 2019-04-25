# Git LFS API Proxy
LFSリクエストを解釈し，ストレージバックエンドのminioからPreSignedURLを生成して返却するAPIプロキシ
- [Git LFS Batch API](https://github.com/git-lfs/git-lfs/blob/master/docs/api/batch.md) の Upload と Download のみ実装
- 認証なし，FileLocking未対応

## 使い方
### サーバ側
`config.json`と`docker-compose.prod.yml`を適宜修正する
```
% docker-compose -f docker-compose.prod.yml up --build
```

### クライアント側
`.lfsconfig`を生成してLFSサーバを指定する
```
% git config -f .lfsconfig remote.origin.lfsurl http://localhost:8080/
```
