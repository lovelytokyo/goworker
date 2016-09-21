# 概要
redisからdequeueして動画をダウンロード、サムネイルを生成する検証を行う。CPU使用率と実行時間を計測する

## 検証環境
mac os x Yosemite memory 16GB

## 実行手順

- 必要なディレクトリー作成

    ```
    cd /path/to/this
    mkdir ./img
    mkdir ./video
    ```
    
- redis起動

    ```
    redis-server /usr/local/etc/redis.conf
    ```
    
- worker起動

    ```
    go run worker.go -queues=default
    ```
    
- enqueue

    ```
    sh push.sh 10
    ```

## 内容

- generator
    - ビデオファイルからサムネイルを生成するpython
    
    ```
    Usage:
      ./generator <video> <interval> <width> <height> <columns> <output>
      ./generator (-h | --help)
      ./generator --version
    
    Options:
      -h --help     Show this screen.
      --version     Show version.
      <video>         Video filepath.
      <interval>      Interval em seconds between frames.
      <width>         Width of each thumbnail.
      <height>        Height of each thumbnail.
      <columns>       Total number of thumbnails per line.
      <output>        Output.
    ```
- push.sh
    - 1から指定個数分、IDを変えながらredisにrpushする
    
    ```
    $ sh push.sh 10
    
    {"class":"MyClass","args":["a123456", "z1", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}
    {"class":"MyClass","args":["a123456", "z2", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}
    {"class":"MyClass","args":["a123456", "z3", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}
    ```

- worker.go
    - redisからdequeue、ビデオをダウンロードしサムネイル生成するpythonプログラムを実行する
        - `generator ./video/{creative_id}.mpeg 2 126 73 4 ./img/{creative_id}.jpg`
    - args
        - [0]: ad_group_id
        - [1]: creative_id
        - [2]: video_url
    
## 検証結果
 - 7プロセス→CPU使用率60%くらい、26プロセス→CPU使用率80%くらい
 - 300本処理に16分 
 　→1本辺り、3.2秒

