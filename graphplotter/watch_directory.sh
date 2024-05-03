#!/bin/bash

# 監視するディレクトリパス
WATCH_DIR="/app/scripts"

# 監視対象のイベント（新しいファイルが作成されたとき）
EVENT="CREATE"

# inotifywaitを使用してディレクトリを監視し、ファイルが追加されたら指定されたコマンドを実行
inotifywait -m -e "$EVENT" --format "%w%f" "$WATCH_DIR" | while read FILE
do
    echo "New file added: $FILE"
    gnuplot $FILE >> /app/scripts/logfile.log 2>> /app/scripts/error.log
    echo "Graph plotted: $FILE"
done
