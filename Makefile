# ----- Makefile by ptr -----
# このファイルをMakefileという名前でカレントディレクトリに保存する
# 使い方:
#    "make test"で課題をすべてテストする。
#     SSH鍵をサーバーに登録しておくと、"make"でローカルのmain.goをサーバーにコピーできる。

# ！ここを変更する！
# 自分に割り当てられたポート番号
PORT=4000
# 自分のtraQ ID
ID=ras
# !ここまで!


# ！ここから（分からない人は）変更しない！
all:
	scp main.go naro:~/go/src/hello-server/main.go

test:
	@echo "\n====================\n"
	@echo "[TEST] /$(ID)"
	curl -X GET "http://localhost:$(PORT)/$(ID)"
	@echo "\n====================\n"
	@echo "[TEST] /ping"
	curl -X GET "http://localhost:$(PORT)/ping"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 1of2"
	curl -X GET "http://localhost:$(PORT)/fizzbuzz?count=20"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 2of2"
	curl -X GET "http://localhost:$(PORT)/fizzbuzz"
	@echo "\n====================\n"
	@echo "[TEST] /add"
	curl -X POST "http://localhost:$(PORT)/add" -d "left=18781&right=18783"
	@echo "\n====================\n"
	@echo "[TEST] /students"
	curl -X GET "http://localhost:$(PORT)/students/3/1"