# MyTools
hasjing的个人工具集
## bli_rename
	bli_rename  go编写的针对哔哩哔哩网站下载的视频进行批量命名的工具
## FolderCompare v1.0
	FolderCompare 基于源目录与目标目录中的文件进行比较，给出源目录中不同的文件清单
		文件的比较算法基于MD5或SHA算法；
	 用法：FolderCompare -t=MD5|SHA -s=源目录 -d=目标目录 -o=输出差异清单文件名
		  -d string
			比较的目标目录（基线目录）,该参数必须设定
		  -h    查看帮助信息
		  -o string
			设置文件差异清单输出的文件，默认输出到屏幕
		  -s string
			比较的源目录 默认为当前目录 (default ".")
		  -t string
			设置比较算法 MD5 或 SHA  默认为MD5 (default "MD5")
