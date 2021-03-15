# youtubeいいね通知
youtubeのいいね数を、一日一回discordに通知する

<img src="img/goodcount.png" alt="goodcount" title="goodcount">

<br>

1) youtubeで動画をいいねする。
2) IFTTTを使って、golangで作ったapiのエンドポイントにpostリクエストをする。
3) 指定の時間になったらdiscordのwebhookを叩いて、いいねした動画を通知する。  

<br>

<img src="img/good.png" alt="goodcount" title="goodcount">
