#!/usr/bin/env sh
go run main.go http://www.animeanime.jp/ http://www.animeanime.biz/ http://www.animate.tv/ http://www.animetopics.com/ http://anime.webnt.jp/ http://anime.excite.co.jp/ http://natalie.mu/comic http://ews.dengeki.com/ http://b.hatena.ne.jp/hotentry/game http://www.famitsu.com/anime/ http://www.presepe.jp/anime_news http://mantan-web.jp/

#結果
#Get "http://anime.excite.co.jp/": dial tcp: lookup anime.excite.co.jp: no such host
#0.66s   93164 http://natalie.mu/comic
#0.66s  102551 http://www.animeanime.jp/
#0.73s  144570 http://news.dengeki.com/
#0.76s    1558 http://www.animate.tv/
#1.04s   46635 http://mantan-web.jp/
#Get "http://www.animetopics.com/": dial tcp: lookup www.animetopics.com: no such host
#1.13s  110306 http://www.animeanime.biz/
#1.14s  398420 http://b.hatena.ne.jp/hotentry/game
#1.23s  228287 http://www.famitsu.com/anime/
#Get "http://www.presepe.jp/anime_news": dial tcp 218.219.148.111:80: i/o timeout
#Get "http://anime.webnt.jp/": dial tcp 54.248.96.159:80: i/o timeout
#30.01s elapsed
