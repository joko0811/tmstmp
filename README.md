# Introduction
作業量全然ないけど何かやった気持ちになってしまうのが嫌で1日の行動ログを取るようにしているのですが、毎回時間を手打ちするのが面倒なので支援ツールを作成しました。
作業内容を現在時間と一緒に記録して
時間と行動内容を指定したファイルに記載するだけの簡単なツールです
# Install

```
$ git clone https://github.com/joko0811/timestamp
$ go mod tidy
$ make build
```

# Setup
前作業　：　ビルドしてできたバイナリにパスを通してください


```
$ tmstmp init

{設定ファイルを編集(デフォルトでは./.tmstmp/config)}
```

# Usage

```
$ ./tmstmp -m hogehuga
$ ./tmstmp show
[hh:mm] hogehuga
```
```
$ ./tmstmp -m kakukakusikazika
$ cat ./timestamp.txt
this is default log file

$ ./tmstmp export
$ cat ./timestamp.txt
this is default log file

[hh:mm] hogehuga
[hh:mm] kakukakusikazika
```
