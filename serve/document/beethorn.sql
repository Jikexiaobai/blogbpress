/*
Navicat MySQL Data Transfer

Source Server         : eforinaj
Source Server Version : 50725
Source Host           : localhost:3306
Source Database       : fiber

Target Server Type    : MYSQL
Target Server Version : 50725
File Encoding         : 65001

Date: 2022-03-31 23:48:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_answer
-- ----------------------------
DROP TABLE IF EXISTS `sys_answer`;
CREATE TABLE `sys_answer` (
  `answer_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `topic_id` bigint(20) DEFAULT NULL,
  `doc` text COMMENT '文档下载',
  `content` varchar(255) DEFAULT NULL,
  `is_adoption` tinyint(4) DEFAULT '1' COMMENT '是否被采纳，1未采纳，2已采纳',
  `price` decimal(10,2) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部, 1待审核，2已发布，3拒绝，4草稿',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`answer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_answer
-- ----------------------------
INSERT INTO `sys_answer` VALUES ('28', '1', '27', '', '撒旦发射点发撒地方', '1', '0.00', '0', '0', '', '2', '2022-03-15 12:49:24', '2022-03-15 12:49:24', null);
INSERT INTO `sys_answer` VALUES ('29', '1', '27', '', '的事发生的故事知道', '0', '0.00', '0', '0', '', '2', '2022-03-31 00:33:35', '2022-03-31 00:33:35', null);
INSERT INTO `sys_answer` VALUES ('30', '1', '27', '', '大发噶收到法国撒旦发个', '0', '0.00', '0', '0', '', '2', '2022-03-31 00:33:38', '2022-03-31 00:33:38', null);
INSERT INTO `sys_answer` VALUES ('31', '1', '27', '', '的方式股市大幅改善大哥', '0', '0.00', '0', '0', '', '2', '2022-03-31 00:33:42', '2022-03-31 00:33:42', null);
INSERT INTO `sys_answer` VALUES ('32', '1', '27', '', '豆腐干岁的法国撒地方', '0', '0.00', '0', '0', '', '2', '2022-03-31 00:33:44', '2022-03-31 00:33:44', null);
INSERT INTO `sys_answer` VALUES ('33', '1', '27', '', 'fg士大夫敢死队风格还是的风格和对方给好的', '0', '0.00', '0', '0', '', '2', '2022-03-31 00:33:47', '2022-03-31 00:33:47', null);
INSERT INTO `sys_answer` VALUES ('34', '1', '27', '', '豆腐干岁的法国大根深蒂固岁的法国士大夫', '2', '0.00', '0', '0', '', '2', '2022-03-31 00:33:56', '2022-03-31 00:33:56', null);
INSERT INTO `sys_answer` VALUES ('35', '1', '27', '', '地方很多法国还是发给', '2', '0.00', '0', '0', '', '2', '2022-03-31 00:35:02', '2022-03-31 00:35:02', null);
INSERT INTO `sys_answer` VALUES ('36', '1', '27', '', '大哥的身份噶士大夫公司', '2', '0.00', '0', '0', '', '2', '2022-03-31 00:36:06', '2022-03-31 00:36:06', null);

-- ----------------------------
-- Table structure for sys_article
-- ----------------------------
DROP TABLE IF EXISTS `sys_article`;
CREATE TABLE `sys_article` (
  `article_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '发布的用户',
  `cate_id` bigint(20) DEFAULT NULL COMMENT '分类id',
  `hots` bigint(20) DEFAULT NULL,
  `favorites` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL COMMENT '阅读量',
  `title` varchar(150) DEFAULT NULL COMMENT '标题',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `content` mediumtext COMMENT '内容',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_article
-- ----------------------------
INSERT INTO `sys_article` VALUES ('1', '1', '1', '123', '12312', '3123', '157', '氪讯基于宝塔运行环境安装教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2>服务器要求</h2>\n<p>建议使用<span style=\"color: #e03e2d;\">centos7.6</span>版本，配置要求<span style=\"color: #e03e2d;\">2h2g</span>或<span style=\"color: #e03e2d;\">2h2g</span>以上。</p>\n<hr />\n<ul>\n<li>在安全开放运行端口分别为：<span style=\"color: #e03e2d;\"><strong>3000，8000，8199，8080</strong></span></li>\n</ul>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2sybuk034sfh05gx.png\" /></p>\n<ul>\n<li>安装环境运行插件：<span style=\"color: #e03e2d;\"><strong>Mysql，Nginx，Redis，Pm2，Supervisor</strong></span></li>\n</ul>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2szza3jarxmki9b8.png\" /></p>\n<p><strong><span style=\"color: #e03e2d;\">注意 ：Mysql版本要求5.6以上（推荐5.7），Nginx版本要求1.12以上（推荐1.18），redis版本默认6.2，pm2内置的node版本需要修改为16.13.1。</span></strong></p>\n<ul>\n<li>配置redis权限密码。</li>\n</ul>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2t30bsyq9dxdh9mw.png\" /></p>\n<p>请在redis设置的配置文件里面找到requirepass foobared 复制然后再下面一行粘贴，后面写入自己的redis密码，<strong><span style=\"color: #e03e2d;\">注意这个密码一定要记住而且保密好</span></strong>，然后保存点击服务重载配置即可</p>\n<hr />\n<p>以上就是氪讯宝塔运行环境安装的教程，如有不懂可联系作者，或者加群。</p>', '/public/uploads/2022-01-11/ch2vjx21fvi81shgaz.png', '2', '2021-11-19 21:38:44', '2022-01-11 21:22:00', null, '通过');
INSERT INTO `sys_article` VALUES ('2', '1', '1', '5234', '231', '312', '423', '氪讯基于宝塔静态网站搭建教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<hr />\n<p>添加站点</p>\n<p>域名：<strong><span style=\"color: #e03e2d;\">xxxx.com，www.xxxx.com</span></strong>，<strong><span style=\"color: #e03e2d;\">m.xxxx.com</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意m开头为h5前台程序域名 ，并且数据库需要选择utf8mb4,如下图</span></strong></p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2ucmysz0d05ynppp.png\" /></p>\n<p>以上就是宝塔创建静态网站教程，如有不懂可联系作者，或者加群。</p>', '/public/uploads/2022-01-11/ch2vjr5jd1010skd6f.png', '2', '2021-12-15 22:33:31', '2022-01-11 21:45:46', null, '通过');
INSERT INTO `sys_article` VALUES ('3', '1', '1', '2131', '23124', '12314', '688', '氪讯基于宝塔数据库导入教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<hr />\n<p>选择宝塔对应创建的数据库，点击导入上传程序包里面的krxun.sql文件然后点击导入即可</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2vxwf7zb8lbkj2ng.png\" /></p>\n<p>是否导入成功，点击数据库右边工具按钮，看到表列表即为导入成功</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2vypdzlqdfu2pc09.png\" /></p>\n<p>以上就是氪讯基于宝塔数据库导入教程，如有不懂可联系作者，或者加群。</p>\n<p>&nbsp;</p>', '/public/uploads/2022-01-11/ch2vkuzuaykxpjtphi.png', '2', '2021-08-24 19:45:25', '2022-01-11 21:46:15', null, '通过');
INSERT INTO `sys_article` VALUES ('4', '1', '1', '3123', '123123', '123125', '123223', '氪讯基于宝塔安装Golang接口服务教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<div class=\"index-module_textWrap_3ygOc\">\n<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">还有必须先把《氪讯基于宝塔数据库导入教程》这一步做完才可进行下面的安装教程</span></strong></p>\n<hr />\n<p>进入网站根目录，创建名为：serve的文件夹。如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2uuqhe1qposjsmd7.png\" /></p>\n<p>进入serve文件夹，分别上传：config，log，public，krxun接口服务程序和文件夹如下图（建议压缩成zip再上传方便点）</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2v3uxklpg5n6wfsr.png\" /></p>\n<p>进入config文件夹打开config.toml文件，进行服务配置</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2vacynb867yq5p5a.png\" /></p>\n<p>数据库配置：mysql:数据库用户名:数据库密码@tcp(127.0.0.1:3306)/数据库名</p>\n<p>mysql:krxun_com:SXEnytcfdfyk@tcp(127.0.0.1:3306)/krxun_com</p>\n<p>Redis配置：127.0.0.1:6379,redis数据库,redis连接密码?idleTimeout=600</p>\n<p>127.0.0.1:6379,2,eforsa_rass_em?idleTimeout=600</p>\n<p>token密钥配置：随机32位字符串仅允许字母数字</p>\n<p>IsDemo配置设置：false</p>\n<p>设置完成保存。</p>\n<hr />\n<p>去到软件商城找到安装的Supervisor管理器，点击设置添加后台守护进程如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2vhlk7hoq5u8ggx4.png\" /></p>\n<p>名称可以随便起。</p>\n<p>运行目录：你的网站根目录下的serve目录</p>\n<p>启动命令：你的网站根目录下的serve目录下的接口服务程序，这里再serve/接口程序名称即可</p>\n<hr />\n<p>测试是否安装成功打开浏览器访问地址:http://你的服务器外网ip地址:8199/api/v1/web/system/kr，如果出现下图表示服务启动成功。</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2w9pdm6tfllryf36.png\" /></p>\n<p>以上就是氪讯接口服务程序安装教程，如有问题可联系作者。</p>\n</div>', '/public/uploads/2022-01-11/ch2vjm25z63wt8j2g7.png', '2', '2021-06-15 23:43:46', '2022-01-11 21:55:01', null, '通过');
INSERT INTO `sys_article` VALUES ('5', '1', '1', '123', '122', '124', '2243', '氪讯基于宝塔安装web前台教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">还有必须先把《氪讯基于宝塔安装Golang接口服务教程》这一步做完才可进行下面的安装教程</span></strong></p>\n<hr />\n<p>进入网站根目录，创建名为：web的文件夹。如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wf6bx6h1ka7ohit.png\" /></p>\n<p>进入web文件夹，分别上传：.nuxt，static，.env，node_modules.zip，nuxt.config.js，package.json文件夹和文件夹（建议压缩成zip再上传方便点），上传后并且解压node_modules.zip到当前目录如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wh8lxhhsxbzjjmf.png\" /></p>\n<hr />\n<p>去到软件商城找到安装的Pm2管理器，点击设置添加web前台守护进程如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wjpcpjzazd0ch8k.png\" /></p>\n<p>项目名称可以随便起。</p>\n<p>运行目录：你的网站根目录下的web目录下的.nuxt</p>\n<p>启动命令：npm run start</p>\n<p>注意这里的node版本一定得是16.13.1。</p>\n<p>浏览器输入测试是否启动成功：http://你的服务器ip地址:3000/</p>\n<p>以上就是氪讯基于宝塔安装web前台教程，如有问题可联系作者。</p>', '/public/uploads/2022-01-11/ch2wmmvmf9kf40qkgj.png', '2', '2021-07-10 23:59:59', '2022-01-11 22:13:05', null, '通过');
INSERT INTO `sys_article` VALUES ('6', '1', '1', '1123', '121', '1213', '3422', '氪讯基于宝塔安装Admin后台管理教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">还有必须先把《氪讯基于宝塔安装Golang接口服务教程》这一步做完才可进行下面的安装教程</span></strong></p>\n<hr />\n<p>进入网站根目录，创建名为：admin的文件夹。如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wogbwr0vil77kxk.png\" /></p>\n<p>进入admin文件夹，上传程序包admin文件夹内的所有文件和文件夹如下图（建议压缩成zip再上传方便点）</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wqirnfzdttiuuuh.png\" /></p>\n<p>上传完成即可，无需配置。</p>\n<p>测试是否安装成功打开浏览器访问地址:http://你的服务器外网ip地址:8000，如果出现下图表示服务启动成功。</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2wsd9qg0k0uldjks.png\" /></p>\n<p>以上就是氪讯基于宝塔安装Admin后台管理教程，如有问题可联系作者。</p>', 'http://krxun.com/public/uploads/2022-01-11/ch2wsoxfpd3lrhubeh.png', '2', '2022-01-11 22:22:38', '2022-01-11 22:22:38', null, '');
INSERT INTO `sys_article` VALUES ('7', '1', '1', '0', '0', '0', '0', '氪讯基于宝塔静态网站配置教程', '氪讯是基于Golang + Vue 开发的前后端分离轻社交内容管理系统，能够解决您快速搭建垂直内容社区。', '<h2 id=\"wznav_1\">推荐的运行环境</h2>\n<p><strong>服务器要求：<span style=\"color: #e03e2d;\">2h2g以上或2h2g</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">注意一定要放行4个端口分别为：3000，8000，8199，8080</span></strong></p>\n<p><strong><span style=\"color: #e03e2d;\">还有必须先把《氪讯基于宝塔安装Golang接口服务教程》这一步做完才可进行下面的安装教程</span></strong></p>\n<hr />\n<p>打开程序包的config文件夹，里面的nginx.conf使用编辑器打开，然后修改里面对应的域名内容与端口内容，后点击你创建对应的网站点击设置按钮，粘贴修改后的nginx.conf内容到配置文件内，如下图</p>\n<p><img src=\"http://www.krxun.com/public/uploads/2022-01-11/ch2x1orivpb86qabnq.png\" /></p>\n<p>保存，然后浏览器输入域名访问测试如果正常表示安装成功。如有问题，就是里面内容配置错误了，实在不懂配置可以找作者付费帮忙。。</p>', '/public/uploads/2022-01-11/ch2x4jbzh4zdfpdyx5.png', '2', '2022-01-11 22:35:19', '2022-01-11 22:35:19', null, '通过');

-- ----------------------------
-- Table structure for sys_audio
-- ----------------------------
DROP TABLE IF EXISTS `sys_audio`;
CREATE TABLE `sys_audio` (
  `audio_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `cate_id` bigint(20) DEFAULT NULL COMMENT '分类',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `link` varchar(255) DEFAULT NULL COMMENT '视频地址',
  `hots` bigint(20) DEFAULT NULL COMMENT '热度',
  `likes` bigint(20) DEFAULT NULL COMMENT '点赞数',
  `favorites` bigint(20) unsigned DEFAULT NULL COMMENT '收藏',
  `views` bigint(20) DEFAULT NULL COMMENT '播放量',
  `has_down` tinyint(4) DEFAULT NULL COMMENT '是否有下载1没有，2有',
  `down_mode` tinyint(4) DEFAULT NULL COMMENT '下载权限 0公开下载，1付费下载，2评论下载，3登录下载',
  `price` decimal(10,2) DEFAULT NULL,
  `down_url` text,
  `purpose` text,
  `attribute` text,
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `other_link` varchar(255) DEFAULT NULL COMMENT '第三方地址',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `delete_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`audio_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_audio
-- ----------------------------
INSERT INTO `sys_audio` VALUES ('2', '1', '4', '浩室人声采样音色Beat包Apollo Sound Authentic House Vocals', 'http://localhost:8199/public/uploads/2021-11-19/cftny9chbl1amnxtzl.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '2131', '123', '23123', '216', '2', '1', '20.00', '[{\"key\":\"baidu.com\",\"title\":\"百度网盘\",\"val\":\"xxsa\"}]', '[{\"key\":\"使用授权\",\"val\":\"商业演出\"}]', '[{\"key\":\"bpm\",\"val\":\"128\"}]', '适当放大风格士大夫感到风格豆腐干', '', '2', null, '2021-07-12 00:42:51', '2021-11-19 17:54:30', '');
INSERT INTO `sys_audio` VALUES ('3', '1', '4', '休闲旋律吉他采样Beat音色包Apollo Sound Chill-Tempo Acoustic Guitar', 'http://localhost:8199/public/uploads/2021-11-19/cftnvggid63b6lts3p.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '23123', '21312', '1231', '3187', '1', '0', '0.00', '[{\"key\":\"baidu.com\",\"title\":\"百度王地方各级\",\"val\":\"213\"}]', '[]', '[]', 'Chill-Tempo Acoustic Guitar是一套吉他采样BEAT音色包，这套音色是由著名的厂牌Apollo Sound 精心制作的音色，包很适合做流行RNB等的风格，非常有感觉的吉他旋律和节奏型，非常不错的一套采样包，适用性非常强，无论你做什么风格的.', '', '2', null, '2021-08-24 20:25:20', '2021-11-19 17:52:52', '');
INSERT INTO `sys_audio` VALUES ('4', '1', '4', '2021流行嘻哈Trap采样鼓组音色Beat包Deedotwill 2021 Drum Kit', 'http://localhost:8199/public/uploads/2021-11-19/cfto04fi4z4hqz0iea.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '123', '312312', '12312', '3135', '1', '0', '0.00', '[]', '[]', '[]', '这是一个来自美洲原著民采样包，由当地的原著民录制，带有手鼓和长笛的当地民族乐器和原生态的民族人声。非常棒的一套采样包，可以让你一下子有回归大自然的感觉，生生不息！', '', '2', null, '2021-11-19 17:56:38', '2021-11-19 17:56:38', '');
INSERT INTO `sys_audio` VALUES ('5', '1', '4', 'Tech House浩室元素采样音色包Toolroom Housewerk', 'http://localhost:8199/public/uploads/2021-11-19/cfto13hrsbm66lshhl.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '2312', '123', '3123', '123131', '1', '0', '0.00', '[]', '[]', '[]', 'Housewerk是一套带有浩室舞曲的采样BEAT套装，非常出色的Tech House浩室风格的元素，如果你喜欢Tech House类的，那么这套采样是比较值得收藏的BEATS精选！无理由5星推荐', '', '2', null, '2021-11-19 17:57:49', '2021-11-19 17:57:49', '');
INSERT INTO `sys_audio` VALUES ('6', '1', '4', '未来流行-鬼魂MIDI工程综合采样音色包W.A Production Future Pop Ghost WAV MiDi', 'http://localhost:8199/public/uploads/2021-11-19/cfto1rq6w4rym2ov9t.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '3123123', '123123', '312', '12333', '1', '0', '0.00', '[]', '[]', '[]', 'Future Pop Ghost是一套带有midi的综合采样BEAT套装，电音风格里非常主流的未来低音和流行风格的采样包，无论是MIDI和预设都是很好的让你快速度了解专业的制作人的制作灵感和参考', '', '2', null, '2021-11-19 17:58:47', '2021-11-19 17:58:47', '');
INSERT INTO `sys_audio` VALUES ('7', '1', '4', '未来低音Future Bass-阴影MIDI工程综合采样音色包W.A Production Future Bass Shadow WAV MiDi SERUM SYLENTH', 'http://localhost:8199/public/uploads/2021-11-19/cfto2qrr3d25zm0lr9.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '123', '12123', '123123', '12314', '1', '0', '0.00', '[]', '[]', '[]', '在trance没落而众多新的风格出现后，还在坚持玩trance的制作人可以试试在Psytrance上找出路，而这个套装里面含有非常出色的音色和亮点，如果你还在玩trance，那么我建议你收藏一下这套Psytrance World的BEATS！', '', '2', null, '2021-11-19 18:00:06', '2021-11-19 18:00:06', '');

-- ----------------------------
-- Table structure for sys_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority`;
CREATE TABLE `sys_authority` (
  `authority_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `title` varchar(50) NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `component` varchar(100) DEFAULT NULL COMMENT '组件地址',
  `order_by` int(10) DEFAULT NULL,
  `redirect` varchar(200) DEFAULT NULL COMMENT '跳转',
  `path` varchar(200) DEFAULT '#' COMMENT '请求地址',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `target` tinyint(4) DEFAULT '1' COMMENT '打开方式（1页签 2新窗口）',
  `type` tinyint(4) DEFAULT '1' COMMENT '菜单类型（1目录 2菜单 3按钮）',
  `hidden` tinyint(4) DEFAULT '2' COMMENT '菜单状态（2显示 1隐藏）',
  `icon` varchar(100) DEFAULT '#' COMMENT '菜单图标',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=164 DEFAULT CHARSET=utf8 COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_authority
-- ----------------------------
INSERT INTO `sys_authority` VALUES ('1', '权限管理', '0', 'PageView', '1', '/authority/list', '/authority', '/authority', '1', '1', '2', 'bars', '2021-09-18 23:49:59', '2021-09-20 13:47:51', '');
INSERT INTO `sys_authority` VALUES ('2', '权限列表', '1', 'AuthorityList', null, null, '/authority/list', '/authority/list', '1', '2', '2', '', '2021-09-18 23:51:48', '2021-09-18 23:51:50', '');
INSERT INTO `sys_authority` VALUES ('3', '权限创建', '1', 'AuthorityCreate', null, null, '/authority/create', '/authority/create', '1', '2', '1', '#', '2021-09-19 17:34:52', '2021-09-19 17:34:54', '');
INSERT INTO `sys_authority` VALUES ('4', '权限编辑', '1', 'AuthorityEdit', '0', '', '/authority/edit', '/authority/edit', '1', '2', '1', '', null, '2021-09-19 21:41:28', '');
INSERT INTO `sys_authority` VALUES ('5', '权限删除', '1', '', '0', '', '', '/authority/remove', '0', '3', '0', '', '2021-09-19 21:43:17', '2021-09-19 21:43:17', '');
INSERT INTO `sys_authority` VALUES ('8', '管理列表', '1', 'MangerList', '2', '', '/manger/list', '/manger/list', '1', '2', '2', '', '2021-09-19 22:48:29', '2021-09-19 22:48:29', '');
INSERT INTO `sys_authority` VALUES ('9', '管理创建', '1', 'MangerCreate', '0', '', '/manger/create', '/manger/create', '1', '2', '1', '', '2021-09-19 23:05:14', '2021-09-20 19:41:27', '');
INSERT INTO `sys_authority` VALUES ('10', '角色列表', '1', 'RoleList', '3', '', '/role/list', '/role/list', '1', '2', '2', '', '2021-09-19 23:10:25', '2021-09-19 23:10:25', '');
INSERT INTO `sys_authority` VALUES ('11', '角色创建', '1', 'RoleCreate', '0', '', '/role/create', '/role/create', '1', '2', '1', '', null, '2021-09-20 00:27:31', '');
INSERT INTO `sys_authority` VALUES ('12', '角色编辑', '1', 'RoleEdit', '0', '', '/role/edit', '/role/edit', '1', '2', '1', '', null, '2021-09-20 01:16:42', '');
INSERT INTO `sys_authority` VALUES ('14', '管理编辑', '1', 'MangerEdit', '0', '', '/manger/edit', '/manger/edit', '1', '2', '1', '', '2021-09-20 13:09:29', '2021-09-20 19:41:17', '');
INSERT INTO `sys_authority` VALUES ('15', '管理删除', '1', '', '0', '', '', '/manger/remove', '0', '3', '0', '', '2021-09-20 13:34:09', '2021-09-20 19:41:05', '');
INSERT INTO `sys_authority` VALUES ('16', '系统管理', '0', 'PageView', '2', '/system/setting', '/system', '/system', '1', '1', '2', 'bars', '2021-09-20 13:46:52', '2021-09-20 13:47:58', '');
INSERT INTO `sys_authority` VALUES ('17', '系统设置', '16', 'SystemSetting', '10', '', '/system/setting', '/system/setting', '1', '2', '2', '', '2021-09-20 13:51:45', '2021-09-20 13:51:45', '');
INSERT INTO `sys_authority` VALUES ('18', '系统保存', '16', '', '0', '', '', '/system/save', '0', '3', '0', '', '2021-09-20 14:02:06', '2021-09-20 14:02:06', '');
INSERT INTO `sys_authority` VALUES ('19', '支付配置', '16', 'PayOption', '2', '', '/pay/option', '/pay/option', '1', '2', '2', '', '2021-09-20 16:19:18', '2021-09-20 16:19:18', '');
INSERT INTO `sys_authority` VALUES ('20', '存储配置', '16', 'OssOption', '0', '', '/oss/option', '/oss/option', '1', '2', '2', '', '2021-09-20 16:30:04', '2021-09-20 16:30:11', '');
INSERT INTO `sys_authority` VALUES ('21', '邮箱配置', '16', 'EmailOption', '0', '', '/email/option', '/email/option', '1', '2', '2', '', '2021-09-20 16:30:40', '2021-09-20 16:30:45', '');
INSERT INTO `sys_authority` VALUES ('27', '系统工具', '0', 'PageView', '3', '/media/list', '/aggregation', '/aggregation', '1', '1', '2', 'bars', '2021-09-20 20:25:39', '2021-09-22 20:07:42', '');
INSERT INTO `sys_authority` VALUES ('28', '媒体列表', '27', 'MediaList', '1', '', '/media/list', '/media/list', '1', '2', '2', '', '2021-09-20 21:18:33', '2021-09-20 21:18:33', '');
INSERT INTO `sys_authority` VALUES ('29', '文章管理', '0', 'PageView', '6', '/article/list', '/article', '/article', '1', '1', '2', 'bars', '2021-09-20 21:20:18', '2021-09-23 12:09:40', '');
INSERT INTO `sys_authority` VALUES ('30', '文章列表', '29', 'ArticleList', '0', '', '/article/list', '/article/list', '1', '2', '2', '', '2021-09-20 21:21:38', '2021-09-23 00:32:49', '');
INSERT INTO `sys_authority` VALUES ('31', '文章创建', '29', 'ArticleCreate', '0', '', '/article/create', '/article/create', '1', '2', '2', '', '2021-09-20 21:22:18', '2021-09-23 01:00:52', '');
INSERT INTO `sys_authority` VALUES ('32', '媒体删除', '28', '', '0', '', '', '/media/remove', '0', '3', '0', '', '2021-09-22 16:35:17', '2021-09-22 16:35:17', '');
INSERT INTO `sys_authority` VALUES ('34', '媒体上传', '28', '', '0', '', '', '/media/upload', '0', '3', '0', '', '2021-09-22 22:00:29', '2021-09-22 22:00:29', '');
INSERT INTO `sys_authority` VALUES ('35', '文章审核', '30', '', '0', '', '', '/article/review', '0', '3', '0', '', '2021-09-22 22:04:53', '2021-09-22 22:04:53', '');
INSERT INTO `sys_authority` VALUES ('36', '文章编辑', '29', 'ArticleEdit', '0', '', '/article/edit', '/article/edit', '1', '2', '1', '', '2021-09-22 22:05:42', '2021-09-23 00:30:32', '');
INSERT INTO `sys_authority` VALUES ('37', '文章编辑信息', '36', '', '0', '', '', '/article/edit/info', '0', '3', '0', '', '2021-09-22 22:06:09', '2021-09-22 22:06:17', '');
INSERT INTO `sys_authority` VALUES ('38', '文章移入回收站', '30', '', '0', '', '', '/article/recover', '0', '3', '0', '', '2021-09-22 22:07:03', '2021-09-22 22:09:52', '');
INSERT INTO `sys_authority` VALUES ('39', '文章还原', '30', '', '0', '', '', '/article/reduction', '0', '3', '0', '', '2021-09-22 22:07:21', '2021-09-22 22:07:21', '');
INSERT INTO `sys_authority` VALUES ('40', '文章删除', '30', '', '0', '', '', '/article/remove', '0', '3', '0', '', '2021-09-22 22:07:44', '2021-09-22 22:07:44', '');
INSERT INTO `sys_authority` VALUES ('42', '分类管理', '45', 'CategoryList', '0', '', '/category/list', '/category/list', '1', '2', '2', '', '2021-09-23 00:32:25', '2021-09-23 01:03:58', '');
INSERT INTO `sys_authority` VALUES ('43', '分类删除', '42', '', '0', '', '', '/category/remove', '0', '3', '0', '', '2021-09-23 00:48:35', '2021-09-23 00:48:35', '');
INSERT INTO `sys_authority` VALUES ('44', '分类创建', '45', 'CategoryCreate', '0', '', '/category/create', '/category/create', '1', '2', '1', '', '2021-09-23 00:49:03', '2021-10-14 14:05:50', '');
INSERT INTO `sys_authority` VALUES ('45', '运营中心', '0', 'PageView', '5', '', '/operation', '/operation', '1', '1', '2', 'bars', '2021-09-23 01:02:01', '2021-09-23 12:09:35', '');
INSERT INTO `sys_authority` VALUES ('46', '标签列表', '45', 'TagList', '0', '', '/tag/list', '/tag/list', '1', '2', '2', '', '2021-09-23 01:17:22', '2021-09-23 01:17:22', '');
INSERT INTO `sys_authority` VALUES ('47', '标签推荐', '46', '', '0', '', '', '/tag/top', '0', '3', '0', '', '2021-09-23 01:18:59', '2021-09-23 01:18:59', '');
INSERT INTO `sys_authority` VALUES ('48', '标签删除', '46', '', '0', '', '', '/tag/remove', '0', '3', '0', '', '2021-09-23 01:19:22', '2021-09-23 01:19:22', '');
INSERT INTO `sys_authority` VALUES ('52', '音频管理', '0', 'PageView', '7', '/audio/list', '/audio', '/audio', '1', '1', '2', 'bars', '2021-09-23 12:11:25', '2021-09-23 12:11:25', '');
INSERT INTO `sys_authority` VALUES ('53', '音频列表', '52', 'AudioList', '0', '', '/audio/list', '/audio/list', '1', '2', '2', '', '2021-09-23 12:16:41', '2021-09-23 12:31:11', '');
INSERT INTO `sys_authority` VALUES ('54', '音频创建', '52', 'AudioCreate', '0', '', '/audio/create', '/audio/create', '1', '2', '2', '', '2021-09-23 12:17:10', '2021-09-23 12:17:10', '');
INSERT INTO `sys_authority` VALUES ('55', '音频编辑', '52', 'AudioEdit', '0', '', '/audio/edit', '/audio/edit', '1', '2', '1', '', '2021-09-23 12:17:39', '2021-09-23 12:17:39', '');
INSERT INTO `sys_authority` VALUES ('56', '查询模块分类', '42', '', '0', '', '', '/category/module', '0', '3', '0', '', '2021-09-23 12:29:16', '2021-09-23 12:29:16', '');
INSERT INTO `sys_authority` VALUES ('57', '热门标签', '46', '', '0', '', '', '/tag/hots', '0', '3', '0', '', '2021-09-23 13:09:54', '2021-09-23 13:09:54', '');
INSERT INTO `sys_authority` VALUES ('58', '音频删除', '53', '', '0', '', '', '/audio/remove', '0', '3', '0', '', '2021-09-23 14:35:29', '2021-09-23 14:37:17', '');
INSERT INTO `sys_authority` VALUES ('59', '音频还原', '53', '', '0', '', '', '/audio/reduction', '0', '3', '0', '', '2021-09-23 14:35:52', '2021-09-23 14:37:09', '');
INSERT INTO `sys_authority` VALUES ('60', '音频移入回收站', '53', '', '0', '', '', '/audio/recover', '0', '3', '0', '', '2021-09-23 14:36:13', '2021-09-23 14:36:13', '');
INSERT INTO `sys_authority` VALUES ('61', '音频审核', '53', '', '0', '', '', '/audio/review', '0', '3', '0', '', '2021-09-23 14:36:53', '2021-09-23 14:36:53', '');
INSERT INTO `sys_authority` VALUES ('62', '资源管理', '0', 'PageView', '7', '/resource/list', '/resource', '/resource', '1', '1', '2', 'bars', '2021-09-23 12:11:25', '2021-09-23 12:11:25', '');
INSERT INTO `sys_authority` VALUES ('63', '资源列表', '62', 'ResourceList', '0', '', '/resource/list', '/resource/list', '1', '2', '2', '', '2021-09-23 12:16:41', '2021-09-23 12:31:11', '');
INSERT INTO `sys_authority` VALUES ('64', '资源创建', '62', 'ResourceCreate', '0', '', '/resource/create', '/resource/create', '1', '2', '2', '', '2021-09-23 12:17:10', '2021-09-23 12:17:10', '');
INSERT INTO `sys_authority` VALUES ('65', '资源编辑', '62', 'ResourceEdit', '0', '', '/resource/edit', '/resource/edit', '1', '2', '1', '', '2021-09-23 12:17:39', '2021-09-23 12:17:39', '');
INSERT INTO `sys_authority` VALUES ('66', '资源删除', '63', '', '0', '', '', '/resource/remove', '0', '3', '0', '', '2021-09-23 14:35:29', '2021-09-23 14:37:17', '');
INSERT INTO `sys_authority` VALUES ('67', '资源还原', '63', '', '0', '', '', '/resource/reduction', '0', '3', '0', '', '2021-09-23 14:35:52', '2021-09-23 14:37:09', '');
INSERT INTO `sys_authority` VALUES ('68', '资源移入回收站', '63', '', '0', '', '', '/resource/recover', '0', '3', '0', '', '2021-09-23 14:36:13', '2021-09-23 14:36:13', '');
INSERT INTO `sys_authority` VALUES ('69', '资源审核', '63', '', '0', '', '', '/resource/review', '0', '3', '0', '', '2021-09-23 14:36:53', '2021-09-23 14:36:53', '');
INSERT INTO `sys_authority` VALUES ('70', '视频管理', '0', 'PageView', '7', '/video/list', '/video', '/video', '1', '1', '2', 'bars', '2021-09-23 12:11:25', '2021-09-23 12:11:25', '');
INSERT INTO `sys_authority` VALUES ('71', '视频列表', '70', 'VideoList', '0', '', '/video/list', '/video/list', '1', '2', '2', '', '2021-09-23 12:16:41', '2021-09-23 12:31:11', '');
INSERT INTO `sys_authority` VALUES ('72', '视频创建', '70', 'VideoCreate', '0', '', '/video/create', '/video/create', '1', '2', '2', '', '2021-09-23 12:17:10', '2021-09-23 12:17:10', '');
INSERT INTO `sys_authority` VALUES ('73', '视频编辑', '70', 'VideoEdit', '0', '', '/video/edit', '/video/edit', '1', '2', '1', '', '2021-09-23 12:17:39', '2021-09-23 12:17:39', '');
INSERT INTO `sys_authority` VALUES ('74', '视频删除', '71', '', '0', '', '', '/video/remove', '0', '3', '0', '', '2021-09-23 14:35:29', '2021-09-23 14:37:17', '');
INSERT INTO `sys_authority` VALUES ('75', '视频还原', '71', '', '0', '', '', '/video/reduction', '0', '3', '0', '', '2021-09-23 14:35:52', '2021-09-23 14:37:09', '');
INSERT INTO `sys_authority` VALUES ('76', '视频移入回收站', '71', '', '0', '', '', '/video/recover', '0', '3', '0', '', '2021-09-23 14:36:13', '2021-09-23 14:36:13', '');
INSERT INTO `sys_authority` VALUES ('77', '视频审核', '71', '', '0', '', '', '/video/review', '0', '3', '0', '', '2021-09-23 14:36:53', '2021-09-23 14:36:53', '');
INSERT INTO `sys_authority` VALUES ('78', '课程管理', '0', 'PageView', '7', '/edu/list', '/edu', '/edu', '1', '1', '2', 'bars', '2021-09-23 12:11:25', '2021-09-23 12:11:25', '');
INSERT INTO `sys_authority` VALUES ('79', '课程列表', '78', 'EduList', '0', '', '/edu/list', '/edu/list', '1', '2', '2', '', '2021-09-23 12:16:41', '2021-09-23 12:31:11', '');
INSERT INTO `sys_authority` VALUES ('80', '课程创建', '78', 'EduCreate', '0', '', '/edu/create', '/edu/create', '1', '2', '2', '', '2021-09-23 12:17:10', '2021-09-23 12:17:10', '');
INSERT INTO `sys_authority` VALUES ('81', '课程编辑', '78', 'EduEdit', '0', '', '/edu/edit', '/edu/edit', '1', '2', '1', '', '2021-09-23 12:17:39', '2021-09-23 12:17:39', '');
INSERT INTO `sys_authority` VALUES ('82', '课程删除', '79', '', '0', '', '', '/edu/remove', '0', '3', '0', '', '2021-09-23 14:35:29', '2021-09-23 14:37:17', '');
INSERT INTO `sys_authority` VALUES ('83', '课程还原', '79', '', '0', '', '', '/edu/reduction', '0', '3', '0', '', '2021-09-23 14:35:52', '2021-09-23 14:37:09', '');
INSERT INTO `sys_authority` VALUES ('84', '课程移入回收站', '79', '', '0', '', '', '/edu/recover', '0', '3', '0', '', '2021-09-23 14:36:13', '2021-09-23 14:36:13', '');
INSERT INTO `sys_authority` VALUES ('85', '课程审核', '79', '', '0', '', '', '/edu/review', '0', '3', '0', '', '2021-09-23 14:36:53', '2021-09-23 14:36:53', '');
INSERT INTO `sys_authority` VALUES ('86', '课程报名列表', '78', 'EduJoinList', '0', '', '/edu/joinList', '/edu/joinList', '1', '2', '1', '', '2021-09-24 00:04:00', '2021-09-24 00:04:26', '');
INSERT INTO `sys_authority` VALUES ('87', '社区管理', '0', 'PageView', '5', '/topic/list', '/community', '/community', '1', '1', '2', 'bars', '2021-09-24 00:18:56', '2021-09-24 00:18:56', '');
INSERT INTO `sys_authority` VALUES ('88', '帖子列表', '87', 'TopicList', '0', '', '/topic/list', '/topic/list', '1', '2', '2', '', '2021-09-24 00:19:29', '2021-09-24 00:19:29', '');
INSERT INTO `sys_authority` VALUES ('89', '问题列表', '87', 'QuestionList', '0', '', '/question', '/question/list', '1', '2', '2', '', '2021-09-24 00:27:08', '2021-09-24 00:27:08', '');
INSERT INTO `sys_authority` VALUES ('90', '圈子列表', '87', 'GroupList', '0', '', '/group/list', '/group/list', '1', '2', '2', '', '2021-09-24 00:27:42', '2021-09-24 00:27:42', '');
INSERT INTO `sys_authority` VALUES ('91', '回答列表', '87', 'AnswerList', '0', '', '/answer/list', '/answer/list', '1', '2', '2', '', '2021-09-24 00:34:50', '2021-09-24 00:34:50', '');
INSERT INTO `sys_authority` VALUES ('92', '评论列表', '87', 'CommentList', '0', '', '/comment/list', '/comment/list', '1', '2', '2', '', '2021-09-24 00:35:19', '2021-09-24 00:35:19', '');
INSERT INTO `sys_authority` VALUES ('93', '帖子置顶', '88', '', '0', '', '', '/topic/top', '0', '3', '0', '', '2021-09-26 13:42:20', '2021-09-26 13:42:20', '');
INSERT INTO `sys_authority` VALUES ('94', '帖子审核', '88', '', '0', '', '', '/topic/review', '0', '3', '0', '', '2021-09-26 13:42:45', '2021-09-26 13:42:45', '');
INSERT INTO `sys_authority` VALUES ('95', '帖子回收', '88', '', '0', '', '', '/topic/recover', '0', '3', '0', '', '2021-09-26 13:43:18', '2021-09-26 13:43:18', '');
INSERT INTO `sys_authority` VALUES ('96', '帖子还原', '88', '', '0', '', '', '/topic/reduction', '0', '3', '0', '', '2021-09-26 13:43:46', '2021-09-26 13:43:46', '');
INSERT INTO `sys_authority` VALUES ('97', '帖子删除', '88', '', '0', '', '', '/topic/remove', '0', '3', '0', '', '2021-09-26 13:44:05', '2021-09-26 13:44:14', '');
INSERT INTO `sys_authority` VALUES ('98', '举报列表', '87', 'ReportList', '0', '', '/report/list', '/report/list', '1', '2', '2', '', '2021-09-26 17:17:19', '2021-09-26 17:17:19', '');
INSERT INTO `sys_authority` VALUES ('99', '举报处理', '98', '', '0', '', '', '/report/review', '0', '3', '0', '', '2021-09-26 17:39:12', '2021-09-26 17:39:12', '');
INSERT INTO `sys_authority` VALUES ('100', '举报删除', '98', '', '0', '', '', '/report/remove', '0', '3', '0', '', '2021-09-26 18:40:48', '2021-09-26 18:40:48', '');
INSERT INTO `sys_authority` VALUES ('101', '评论审核', '92', '', '0', '', '', '/comment/review', '0', '3', '0', '', '2021-09-26 18:46:38', '2021-09-26 18:47:45', '');
INSERT INTO `sys_authority` VALUES ('102', '评论移入回收站', '92', '', '0', '', '', '/comment/recover', '0', '3', '0', '', '2021-09-26 18:48:22', '2021-09-26 18:48:22', '');
INSERT INTO `sys_authority` VALUES ('103', '评论还原', '92', '', '0', '', '', '/comment/reduction', '0', '3', '0', '', '2021-09-26 18:48:33', '2021-09-26 18:48:33', '');
INSERT INTO `sys_authority` VALUES ('104', '评论删除', '92', '', '0', '', '', '/comment/remove', '0', '3', '0', '', '2021-09-26 18:48:43', '2021-09-26 18:48:43', '');
INSERT INTO `sys_authority` VALUES ('105', '答案审核', '91', '', '0', '', '', '/answer/review', '0', '3', '0', '', '2021-09-26 18:49:26', '2021-09-26 18:49:26', '');
INSERT INTO `sys_authority` VALUES ('106', '答案移入回收站', '91', '', '0', '', '', '/answer/recover', '0', '3', '0', '', '2021-09-26 18:49:41', '2021-09-26 18:49:41', '');
INSERT INTO `sys_authority` VALUES ('107', '答案还原', '91', '', '0', '', '', '/answer/reduction', '0', '3', '0', '', '2021-09-26 18:49:52', '2021-09-26 18:49:52', '');
INSERT INTO `sys_authority` VALUES ('108', '答案删除', '91', '', '0', '', '', '/answer/remove', '0', '3', '0', '', '2021-09-26 18:50:00', '2021-09-26 18:50:00', '');
INSERT INTO `sys_authority` VALUES ('109', '圈子审核', '90', '', '0', '', '', '/group/review', '0', '3', '0', '', '2021-09-26 18:50:37', '2021-09-26 18:50:37', '');
INSERT INTO `sys_authority` VALUES ('110', '圈子移入回收站', '90', '', '0', '', '', '/group/recover', '0', '3', '0', '', '2021-09-26 18:50:46', '2021-09-26 18:50:46', '');
INSERT INTO `sys_authority` VALUES ('111', '圈子还原', '90', '', '0', '', '', '/group/reduction', '0', '3', '0', '', '2021-09-26 18:50:54', '2021-09-26 18:50:54', '');
INSERT INTO `sys_authority` VALUES ('112', '圈子删除', '90', '', '0', '', '', '/group/remove', '0', '3', '0', '', '2021-09-26 18:51:04', '2021-09-26 18:51:04', '');
INSERT INTO `sys_authority` VALUES ('113', '问题审核', '89', '', '0', '', '', '/question/review', '0', '3', '0', '', '2021-09-26 18:51:39', '2021-09-26 18:51:39', '');
INSERT INTO `sys_authority` VALUES ('114', '问题移入回收站', '89', '', '0', '', '', '/question/recover', '0', '3', '0', '', '2021-09-26 18:51:51', '2021-09-26 18:51:51', '');
INSERT INTO `sys_authority` VALUES ('115', '问题还原', '89', '', '0', '', '', '/question/reduction', '0', '3', '0', '', '2021-09-26 18:52:02', '2021-09-26 18:52:02', '');
INSERT INTO `sys_authority` VALUES ('116', '问题删除', '89', '', '0', '', '', '/question/remove', '0', '3', '0', '', '2021-09-26 18:52:10', '2021-09-26 18:52:10', '');
INSERT INTO `sys_authority` VALUES ('117', '用户管理', '0', 'PageView', '5', '/user/list', '/user', '/user', '1', '1', '2', 'bars', '2021-09-26 21:34:22', '2021-09-26 21:35:55', '');
INSERT INTO `sys_authority` VALUES ('118', '用户列表', '117', 'UserList', '0', '', '/user/list', '/user/list', '1', '2', '2', '', '2021-09-26 21:36:57', '2021-09-26 21:36:57', '');
INSERT INTO `sys_authority` VALUES ('119', '等级列表', '117', 'GradeList', '0', '', '/grade/list', '/grade/list', '1', '2', '2', '', '2021-09-26 21:37:26', '2021-09-26 21:37:26', '');
INSERT INTO `sys_authority` VALUES ('120', '会员列表', '117', 'VipList', '0', '', '/vip/list', '/vip/list', '1', '2', '2', '', '2021-09-26 21:37:43', '2021-09-26 21:37:43', '');
INSERT INTO `sys_authority` VALUES ('121', '实名认证', '117', 'VerifyList', '0', '', '/verify/list', '/verify/list', '1', '2', '2', '', '2021-09-26 21:38:10', '2021-09-26 21:38:10', '');
INSERT INTO `sys_authority` VALUES ('122', '用户创建', '117', 'UserCreate', '0', '', '/user/create', '/user/create', '1', '2', '1', '', '2021-09-27 14:18:37', '2021-09-27 14:18:37', '');
INSERT INTO `sys_authority` VALUES ('123', '用户编辑', '117', 'UserEdit', '0', '', '/user/edit', '/user/edit', '1', '2', '1', '', '2021-09-27 14:19:14', '2021-09-27 14:19:14', '');
INSERT INTO `sys_authority` VALUES ('124', '用户审核', '118', '', '0', '', '', '/user/review', '1', '3', '1', '', '2021-09-27 14:20:08', '2021-09-27 14:20:58', '');
INSERT INTO `sys_authority` VALUES ('125', '用户删除', '118', '', '0', '', '', '/user/remove', '1', '3', '1', '', '2021-09-27 14:20:21', '2021-09-27 14:20:48', '');
INSERT INTO `sys_authority` VALUES ('126', '获取用户修改信息', '123', '', '0', '', '', '/user/edit/info', '0', '3', '0', '', '2021-09-27 16:04:11', '2021-09-27 16:04:11', '');
INSERT INTO `sys_authority` VALUES ('127', '获取课程修改信息', '81', '', '0', '', '', '/edu/edit/info', '0', '3', '0', '', '2021-09-27 16:04:56', '2021-09-27 16:04:56', '');
INSERT INTO `sys_authority` VALUES ('128', '获取视频修改信息', '73', '', '0', '', '', '/video/edit/info', '0', '3', '0', '', '2021-09-27 16:05:16', '2021-09-27 16:05:16', '');
INSERT INTO `sys_authority` VALUES ('129', '获取音频修改信息', '55', '', '0', '', '', '/audio/edit/info', '0', '3', '0', '', '2021-09-27 16:05:33', '2021-09-27 16:05:33', '');
INSERT INTO `sys_authority` VALUES ('130', '获取资源修改信息', '65', '', '0', '', '', '/resource/edit/info', '0', '3', '0', '', '2021-09-27 16:05:50', '2021-09-27 16:05:50', '');
INSERT INTO `sys_authority` VALUES ('131', '等级创建', '119', '', '0', '', '', '/grade/create', '0', '3', '0', '', '2021-09-27 20:59:18', '2021-09-27 20:59:18', '');
INSERT INTO `sys_authority` VALUES ('132', '等级编辑', '119', '', '0', '', '', '/grade/edit', '0', '3', '0', '', '2021-09-27 20:59:33', '2021-09-27 20:59:33', '');
INSERT INTO `sys_authority` VALUES ('133', '会员创建', '120', '', '0', '', '', '/vip/create', '0', '3', '0', '', '2021-09-27 21:39:15', '2021-09-27 21:39:15', '');
INSERT INTO `sys_authority` VALUES ('134', '会员编辑', '120', '', '0', '', '', '/vip/edit', '0', '3', '0', '', '2021-09-27 21:39:30', '2021-09-27 21:39:30', '');
INSERT INTO `sys_authority` VALUES ('135', '获取会员编辑信息', '120', '', '0', '', '', '/vip/edit/info', '0', '3', '0', '', '2021-09-27 21:40:44', '2021-09-27 21:40:44', '');
INSERT INTO `sys_authority` VALUES ('136', '获取等级编辑信息', '119', '', '0', '', '', '/grade/edit/info', '0', '3', '0', '', '2021-09-27 21:40:54', '2021-09-27 21:40:54', '');
INSERT INTO `sys_authority` VALUES ('137', '认证审核', '121', '', '0', '', '', '/verify/review', '0', '3', '0', '', '2021-09-27 21:59:30', '2021-09-27 21:59:30', '');
INSERT INTO `sys_authority` VALUES ('138', '认证删除', '121', '', '0', '', '', '/verify/remove', '0', '3', '0', '', '2021-09-27 21:59:38', '2021-09-27 21:59:38', '');
INSERT INTO `sys_authority` VALUES ('139', '财务管理', '0', 'PageView', '5', '/order/list', '/finance', '/finance', '1', '1', '2', 'bars', '2021-09-27 22:11:12', '2021-09-27 22:11:12', '');
INSERT INTO `sys_authority` VALUES ('140', '订单列表', '139', 'OrderList', '0', '', '/order/list', '/order/list', '1', '2', '2', '', '2021-09-27 22:11:53', '2021-09-27 22:12:53', '');
INSERT INTO `sys_authority` VALUES ('141', '提现列表', '139', 'CashList', '0', '', '/cash/list', '/cash/list', '1', '2', '2', '', '2021-09-27 22:12:42', '2021-09-27 22:12:42', '');
INSERT INTO `sys_authority` VALUES ('142', '订单删除', '140', '', '0', '', '', '/order/remove', '0', '3', '0', '', '2021-09-27 22:48:10', '2021-09-27 22:48:54', '');
INSERT INTO `sys_authority` VALUES ('143', '订单信息查看', '139', 'OrderInfo', '0', '', '/order/info', '/order/info', '1', '2', '1', '', '2021-09-27 22:48:44', '2021-09-27 22:51:10', '');
INSERT INTO `sys_authority` VALUES ('144', '提现审核', '141', '', '0', '', '', '/cash/review', '0', '3', '0', '', '2021-09-27 23:14:35', '2021-11-04 01:24:25', '');
INSERT INTO `sys_authority` VALUES ('145', '提现删除', '141', '', '0', '', '', '/cash/remove', '0', '3', '0', '', '2021-09-27 23:14:45', '2021-09-27 23:14:45', '');
INSERT INTO `sys_authority` VALUES ('146', '获取角色编辑信息', '12', '', '0', '', '', '/role/edit/info', '0', '3', '0', '', '2021-09-28 13:24:32', '2021-09-28 13:24:32', '');
INSERT INTO `sys_authority` VALUES ('147', '获取管理编辑信息', '14', '', '0', '', '', '/manger/edit/info', '0', '3', '0', '', '2021-09-28 13:25:08', '2021-09-28 13:25:08', '');
INSERT INTO `sys_authority` VALUES ('148', '获取权限编辑信息', '4', '', '0', '', '', '/authority/edit/info', '0', '3', '0', '', '2021-09-28 13:25:35', '2021-09-28 13:25:35', '');
INSERT INTO `sys_authority` VALUES ('149', '外观设计', '0', 'PageView', '0', '/design/web', '/design', '/design', '1', '1', '2', 'bars', '2021-10-05 16:22:32', '2021-10-05 16:22:32', '');
INSERT INTO `sys_authority` VALUES ('151', '分类编辑', '45', 'CategoryEdit', '0', '', '/category/edit', '/category/edit', '1', '2', '1', '', '2021-10-14 14:15:24', '2021-10-14 14:15:24', '');
INSERT INTO `sys_authority` VALUES ('152', '获取编辑信息', '42', '', '0', '', '', '/category/edit/info', '0', '3', '0', '', '2021-10-14 14:49:49', '2021-10-14 14:49:49', '');
INSERT INTO `sys_authority` VALUES ('153', '充值列表', '139', 'RechargeList', '0', '', '/recharge/list', '/recharge/list', '1', '2', '2', '', '2021-11-02 23:28:42', '2021-11-02 23:28:42', '');
INSERT INTO `sys_authority` VALUES ('154', '充值删除', '153', '', '0', '', '', '/recharge/remove', '0', '3', '0', '', '2021-11-02 23:29:17', '2021-11-02 23:29:17', '');
INSERT INTO `sys_authority` VALUES ('155', '充值审核', '153', '', '0', '', '', '/recharge/review', '0', '3', '0', '', '2021-11-02 23:29:45', '2021-11-02 23:29:45', '');
INSERT INTO `sys_authority` VALUES ('156', '卡密列表', '139', 'CardList', '0', '', '/card/list', '/card/list', '1', '2', '2', '', '2021-11-04 01:54:26', '2021-11-04 02:08:20', '');
INSERT INTO `sys_authority` VALUES ('157', '生成卡密', '156', '', '0', '', '', '/card/create', '0', '3', '0', '', '2021-11-04 01:54:56', '2021-11-04 02:08:41', '');
INSERT INTO `sys_authority` VALUES ('158', '卡密删除', '156', '', '0', '', '', '/card/remove', '0', '3', '0', '', '2021-11-04 01:55:20', '2021-11-04 02:08:34', '');
INSERT INTO `sys_authority` VALUES ('159', '短信配置', '16', 'SmsOption', '1', '', '/sms/option', '/sms/option', '1', '2', '2', '', '2021-11-28 16:45:51', '2021-11-28 16:45:51', '');
INSERT INTO `sys_authority` VALUES ('160', '首页模块设计', '149', 'DesignHome', '0', '', '/design/home', '/design/home', '1', '2', '2', '', '2021-11-30 10:46:08', '2021-11-30 14:47:54', '');
INSERT INTO `sys_authority` VALUES ('161', '仪表盘', '0', 'Dashboard', '100', '', '/dashboard', '/dashboard', '1', '1', '2', 'bars', '2021-12-07 12:37:46', '2021-12-07 12:49:13', '');
INSERT INTO `sys_authority` VALUES ('162', '大文件上传', '28', '', '0', '', '', '/upload/chunk', '0', '3', '0', '', '2021-12-17 13:37:34', '2021-12-17 13:44:08', '');
INSERT INTO `sys_authority` VALUES ('163', '大文件合并', '28', '', '0', '', '', '/upload/mergeChunk', '0', '3', '0', '', '2021-12-17 16:40:37', '2021-12-17 16:40:37', '');

-- ----------------------------
-- Table structure for sys_card
-- ----------------------------
DROP TABLE IF EXISTS `sys_card`;
CREATE TABLE `sys_card` (
  `card_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `used_id` bigint(20) DEFAULT NULL COMMENT '使用者id',
  `secret_key` varchar(255) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '状态: 1未使用，2已使用',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`card_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_card
-- ----------------------------
INSERT INTO `sys_card` VALUES ('3', '0', '46cd265a093bc72ae77cabec35445d30', '100.00', '1', '2021-11-04 17:10:41');
INSERT INTO `sys_card` VALUES ('4', '0', 'dbbdb3cdbee5fb4c4d400bb25beccf41', '100.00', '1', '2021-11-04 17:10:41');
INSERT INTO `sys_card` VALUES ('5', '0', '71ea00ccfe1217427ac86d2b451256bd', '100.00', '1', '2021-11-04 17:10:41');
INSERT INTO `sys_card` VALUES ('6', '0', '1e09de2ee57421f38571f439f0d14121', '100.00', '1', '2021-11-04 17:10:41');
INSERT INTO `sys_card` VALUES ('7', '0', 'c427dcf02b51b5b93522036a97f3180d', '100.00', '1', '2021-11-04 17:10:41');

-- ----------------------------
-- Table structure for sys_carousel
-- ----------------------------
DROP TABLE IF EXISTS `sys_carousel`;
CREATE TABLE `sys_carousel` (
  `carousel_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mode` tinyint(4) DEFAULT NULL COMMENT '类型1(投稿内容)，2(其他内容)',
  `related_id` bigint(20) DEFAULT NULL,
  `module` varchar(50) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL COMMENT '链接',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面地址',
  `type` tinyint(4) DEFAULT NULL COMMENT '类型: 1(web),2(app)',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`carousel_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_carousel
-- ----------------------------
INSERT INTO `sys_carousel` VALUES ('1', '1', '1', 'article', 'http://localhost:3000', 'http://fibercms.com/public/uploads/2021-08-26/cdtkto2n81q4lcaqgx.jpg', '1', '2021-10-05 10:19:17', null, null);
INSERT INTO `sys_carousel` VALUES ('2', '1', '2', 'audio', 'http://localhost:3000', 'http://fibercms.com/public/uploads/2021-08-26/cdtkb5gyuaf0ytlaha.jpg', '1', '2021-10-05 10:19:43', null, null);
INSERT INTO `sys_carousel` VALUES ('3', '1', '2', 'video', 'http://localhost:3000', 'http://fibercms.com/public/uploads/2021-08-26/cdtjyo0af9dcedko0u.jpg', '1', '2021-10-05 10:46:41', null, null);

-- ----------------------------
-- Table structure for sys_cash
-- ----------------------------
DROP TABLE IF EXISTS `sys_cash`;
CREATE TABLE `sys_cash` (
  `cash_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `code` varchar(50) DEFAULT NULL COMMENT '提现单号',
  `cash_money` decimal(10,2) DEFAULT NULL,
  `service_money` decimal(10,2) DEFAULT NULL COMMENT '服务费',
  `money` decimal(10,2) DEFAULT NULL COMMENT '实际金额',
  `mode` tinyint(4) DEFAULT NULL COMMENT '方式，1人工转账，2第三方到账',
  `number` varchar(255) DEFAULT NULL COMMENT '账号',
  `receipt_num` varchar(255) DEFAULT NULL,
  `pay_method` tinyint(4) DEFAULT NULL COMMENT '支付方式（1支付宝，2微信）',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态 1待审核，2审核通过，3审核不通过',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`cash_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_cash
-- ----------------------------
INSERT INTO `sys_cash` VALUES ('5', '1', 'C1163596113087105514311', '95.00', '5.00', '100.00', '1', 'wrerwe', 'ghfshsdfh', '1', '2', 'sdfgsdfg', '2021-11-04 01:38:50', '2021-11-04 01:38:50');

-- ----------------------------
-- Table structure for sys_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_category`;
CREATE TABLE `sys_category` (
  `cate_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT '0' COMMENT '顶级分类',
  `module` varchar(50) DEFAULT NULL COMMENT '所属模块',
  `title` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `slug` varchar(50) DEFAULT NULL COMMENT '分类别名',
  `cover` varchar(255) DEFAULT NULL COMMENT '分类背景图',
  `sort` int(10) DEFAULT NULL COMMENT '分类排序',
  `keywords` varchar(50) DEFAULT NULL COMMENT '分类关键字',
  `description` varchar(255) DEFAULT NULL COMMENT '分类描述',
  `isTop` tinyint(4) DEFAULT '1' COMMENT '分类状态 1 不推荐 2推荐',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL COMMENT '更新日期',
  `delete_time` datetime DEFAULT NULL COMMENT '删除标志（0代表存在 2代表删除）',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_category
-- ----------------------------
INSERT INTO `sys_category` VALUES ('1', '0', 'article', '商业', 'sy', null, '1', 'sy', 'sy', '2', '2021-09-22 20:58:57', '2021-09-22 20:59:00', null);
INSERT INTO `sys_category` VALUES ('4', '0', 'audio', 'bigroom', 'bg', '', '0', 'bigroom', 'asd', '2', '2021-09-23 13:10:41', '2021-09-23 13:10:41', null);
INSERT INTO `sys_category` VALUES ('5', '0', 'resource', '测试', 'asd', '', '0', '测试', 'sdfsdf', '2', '2021-09-23 16:07:36', '2021-09-23 16:07:36', null);
INSERT INTO `sys_category` VALUES ('6', '0', 'video', '测士大夫', 'asd第三方', '', '0', '测士大夫', '啊撒大苏打', '2', '2021-09-23 17:25:12', '2021-09-23 17:25:12', null);
INSERT INTO `sys_category` VALUES ('7', '0', 'edu', '岁的法国xcgv', 'asddfg', '', '0', '岁的法国', '', '2', '2021-09-23 22:46:17', '2021-09-23 22:46:17', null);
INSERT INTO `sys_category` VALUES ('8', '0', 'article', '科技', 'keji', '', '0', '科技', '科技改变世界', '2', '2021-10-17 18:45:19', '2021-10-17 18:45:19', null);
INSERT INTO `sys_category` VALUES ('9', '0', 'group', '官方', 'guangf', null, null, '啊实打实', '撒旦发射点', '2', '2021-10-28 17:56:23', '2021-10-28 17:56:25', null);

-- ----------------------------
-- Table structure for sys_comment
-- ----------------------------
DROP TABLE IF EXISTS `sys_comment`;
CREATE TABLE `sys_comment` (
  `comment_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `reply_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL COMMENT '发布用户',
  `parent_id` bigint(20) DEFAULT NULL,
  `top_id` bigint(20) DEFAULT NULL,
  `related_id` bigint(20) DEFAULT NULL,
  `module` varchar(50) DEFAULT NULL COMMENT '模块',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `type` tinyint(4) DEFAULT NULL COMMENT '类型 1图片 2视频 3文字',
  `files` text COMMENT '文件链接',
  `likes` bigint(20) DEFAULT NULL COMMENT '点赞数',
  `unlikes` bigint(20) DEFAULT NULL COMMENT '点踩',
  `remark` varchar(500) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '2已审核，1未审核',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=112 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_comment
-- ----------------------------
INSERT INTO `sys_comment` VALUES ('106', '0', '1', '0', '0', '21', 'topic', 's撒大噶高大上的噶啥的噶', '0', '', '0', '0', '', '2', '2022-03-12 17:57:04', '2022-03-12 17:57:04', null);
INSERT INTO `sys_comment` VALUES ('107', '0', '1', '0', '0', '21', 'topic', '大发噶收到法国撒旦发个', '0', '', '0', '0', '', '2', '2022-03-12 17:57:06', '2022-03-12 17:57:06', null);
INSERT INTO `sys_comment` VALUES ('108', '1', '1', '107', '107', '21', 'topic', '手动阀手动阀撒旦', '0', '', '0', '0', '', '2', '2022-03-12 17:57:08', '2022-03-12 17:57:08', null);
INSERT INTO `sys_comment` VALUES ('109', '1', '1', '108', '107', '21', 'topic', '十大发射点发射点发射点', '0', '', '0', '0', '', '2', '2022-03-12 17:57:12', '2022-03-12 17:57:12', null);
INSERT INTO `sys_comment` VALUES ('110', '1', '1', '106', '106', '21', 'topic', '的风格和对方给', '0', '', '0', '0', '', '2', '2022-03-12 19:16:37', '2022-03-12 19:16:37', null);
INSERT INTO `sys_comment` VALUES ('111', '0', '1', '0', '0', '27', 'topic', 'fgdhdsfghdfgh', '0', '', '0', '0', '', '2', '2022-03-15 14:44:53', '2022-03-15 14:44:53', null);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` text COMMENT '参数键值',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES ('5', '邮箱配置', 'EmailOptions', '{\"host\":\"smtp.qq.com\",\"port\":\"587\",\"user\":\"973728679@qq.com\",\"pass\":\"mzrchjdssewlbbib\",\"email\":\"973728679@qq.com\"}', '2020-05-18 00:34:51', '2021-08-17 19:38:37', '邮箱配置');
INSERT INTO `sys_config` VALUES ('7', '基础设置', 'BaseSetting', '{\"title\":\"氪讯\",\"childTitle\":\"轻社区内容管理系统-资源交易-在线网课-社区问答-动态交流\",\"description\":\"氪讯是基于Golang + Vue 开发的前后端分离轻社区内容管理系统，能够解决您快速搭建垂直内容社区。\\n架构：apiSever（go语言编写），后台管理（vue+antdAdmin）,前台用户交互（Nuxtjs+Vue+Antd）\\n模块（目前）：文章，音频，视频，资源，网课，圈子，问答，动态帖子，八大模块功能（未来会加入更多内容模块），能满足各种内容系统刚需模块。\",\"url\":\"http://krxun.com\",\"logo\":\"http://localhost:8199/public/uploads/2021-12-30/cgsmufeakogxomn2wq.png\",\"adminLogo\":\"http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png\",\"icon\":\"http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png\",\"currencySymbol\":\"￥\",\"language\":\"zh\",\"recordNumber\":\"桂-4152055\"}', '2020-05-24 00:41:43', '2022-01-06 14:40:37', '基础设置');
INSERT INTO `sys_config` VALUES ('8', '文件设置', 'FileSetting', '{\"engine\":1,\"path\":\"uploads\",\"fileSize\":\"50\",\"imageType\":[\".png\",\".jpg\",\".jpeg\"],\"audioType\":[\".mp3\"],\"videoType\":[\".mp4\"],\"otherType\":[\".crt\"]}', '2020-05-30 22:47:11', '2021-12-17 20:38:16', '文件设置');
INSERT INTO `sys_config` VALUES ('9', '支付宝配置', 'AlyPayOptions', '{\"appId\":\"2016102700770129\",\"privateKey\":\"MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCjkmnBSEjwpvBFLUF279MLV1bSjjnnokhocdEgTYx4Qy6zhtpUfyV/LwRQ31ddUtoTYhCedOO5nd95vMrAOSU2++ZdXWVjUfw1er4fFaMlNp8hQkKOgp8bDLO05x3/o9rG0oOWqWreyUGoPoWRQhnzJt6kYbJtunx3d8hBN59jM0PHiDODDsgiu8sq4muk/BjayF8FKypkp2LQCNjpxbUPIcbjpLsRzrJ4eH+Om+10cMHpIDLl6JZITwq3JF6l39cS61N4xNwOOtE/1KCo3x9/y8T2XF3PQvfHjXHeNuwQNIvey828a7KFbTgV6O09lOCTvjP1HmA1zvpHD6MizaofAgMBAAECggEAfX9WGG9HVywd6FVihshWGbt257EriCage1Hn62rUPWj+KctrM60hrcT7ALl6pCVvH7P7oDd6iO0xioto27Z8cQUvp66CnYNHiBiWSe8l7uVLjg7yVbiuLei+8CrqfzrOHgFh6HQvhKLQ9y6Q9/PJSR8nbuNuLHYKDPcf87mjVk3Cz6D03V9B1avfaH5IiLvlZwrAgVk7FoZEoBnhZhIg7yME4DUitPykTGg1ur/Gofe+xrPqONOnvZW8wNa/l2nKIPQwBaiUVcSV7gVBRRkw/GADMV59titx1dyzmR0b61GLexjI2WxXRLnAe1+qXrbOcbsCKIgFjQELbCz8MYFYgQKBgQDM7m/rzIcsJkmHXFGZdVRIdFzx73SUKMSu3srCs+7urUG9HNupf16ykMHv1p7rA0tBt6lzTAUgpf0F959JE1fUl+tyy2gUDyqOrP1FzF91wbertTXcrBcv9wMD5mvENSHNz9CDQ3Z7ZHtpPSnk2AvGhXzDSc6pnEkBzUIIVjgQfwKBgQDMVXSsnESsUwzz9qfzWxT9qsELIMM91xyn/w7FWcvIZvd33TzpPP0aZ69hylmaUPpSp4wwL5+zqPOvfI03y5En7ZzohMBTebU4H4m3aK4GHF0w+6ft+bt04ZAyFc7lyf/w+nEniPgQHSQqMu64FbK+GrYxrAXzSFhHfH/b35uWYQKBgGYpK3RSdsRkpd0sAaXN3uFr2PXnGKfPlxVSDaR4jNFBX/dFzp+11mCQV44X4QtpjffJ9lh6+kdnWDbEVgzY7r0VqxOEIXN2iBGuXWiRVLlghA6+fIZw5/JKYp9sHCcpEZwHUHxPgl5LHla9XgguR9iErUixn6vgNGkIiTWcvcBDAoGAfUGNShppBnHKqOp0vfsBfRZlS9sDlC7/RARYG6YWA30LChE2u4tFZCBXJE0UbEJjkLNgflFTRqC08MgbES7ahm1kGCz4cLNU4ViD5UhoFRriDZrWsEy8GsQCzpELyVTwbdo37xJJbidO+gdKytGSRnK9aOmYpC+e3gN1pWUHTUECgYA0zjkcEHfYzxH/VChkT9SzuI4WIJ+NSQ5ztljaZlDPg8rskUCB/kV6tRXW5oVBmgm4VyHl5bOl8PkoUC06Zhe6jiZ5cSUnfvCi7w2ysc1y5Nypc1P/qBu0bGK0nLRdFHbuv4IZsCrcvaR86QFeK14xO+yJ3kkOyXLd4UTvXYG/xg==\",\"appPublicKey\":\"http://localhost:8199/public/uploads/2021-07-07/ccmt7z025mv8n9gz5y.crt\",\"alyCertPublicKey\":\"http://localhost:8199/public/uploads/2021-07-07/ccmt805eqlvknzgjgs.crt\",\"alyRootCert\":\"http://localhost:8199/public/uploads/2021-07-07/ccmt819guk10ixsnmq.crt\"}', '2020-06-13 15:32:17', '2021-08-27 09:43:35', '支付宝配置');
INSERT INTO `sys_config` VALUES ('10', '支付设置', 'PaySetting', '{\"alyPay\":2,\"weChatPay\":2,\"payMode\":[1,3],\"recharge\":[1,3,4],\"cashMin\":50,\"cashServicePercent\":0.05,\"servicePercent\":0.05}', '2020-06-13 16:37:21', '2021-11-11 12:31:59', '支付设置');
INSERT INTO `sys_config` VALUES ('12', '用户设置', 'UserSetting', '{\"defaultAvatar\":\"http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png\",\"defaultCover\":\"http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png\",\"defaultGrade\":\"1\",\"verifyMode\":\"0\",\"verifyPrice\":21}', '2020-10-25 17:25:12', '2021-08-17 13:46:49', '用户设置');
INSERT INTO `sys_config` VALUES ('22', '首页设计', 'HomeDesign', '[{\"title\":\"轮播\",\"height\":\"300\",\"list\":[{\"link\":\"/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftkh4fqsdcgux3vta.png\",\"isPlatform\":2,\"title\":\"测试首页\"},{\"link\":\"https://www.baidu.com/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftsrl7r9nnngyv4ku.jpg\",\"isPlatform\":1,\"title\":\"测试外链\"},{\"link\":\"/edu\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftskpi0mwiskcldfg.jpg\",\"isPlatform\":2,\"title\":\"课程学习\"},{\"link\":\"/feed\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg\",\"isPlatform\":2,\"title\":\"动态\"}],\"style\":1,\"isOpen\":false},{\"title\":\"音频组件\",\"showTitle\":2,\"audioIds\":\"2,3,4,5\",\"style\":5,\"isOpen\":false},{\"title\":\"广告\",\"showTitle\":1,\"height\":\"200\",\"style\":9,\"content\":\"{\\\"cover\\\":\\\"http://localhost:8199/public/uploads/2021-11-19/cftkh4fqsdcgux3vta.png\\\",\\\"link\\\":\\\"/\\\",\\\"isPlatform\\\":2}\",\"isOpen\":false},{\"title\":\"视频组件\",\"showTitle\":2,\"videoIds\":\"2,3,4,5\",\"style\":6,\"isOpen\":false},{\"title\":\"广告2\",\"showTitle\":1,\"height\":\"200\",\"style\":9,\"content\":\"{\\\"cover\\\":\\\"http://localhost:8199/public/uploads/2021-11-19/cftnonwtwzmjcc0hwi.jpg\\\",\\\"link\\\":\\\"/\\\",\\\"isPlatform\\\":2}\",\"isOpen\":false},{\"title\":\"资源组件\",\"showTitle\":2,\"resourceIds\":\"2,3,4,5\",\"style\":7,\"isOpen\":false},{\"title\":\"专栏\",\"showTitle\":2,\"articleIds\":\"1,2,3,4\",\"style\":3,\"isOpen\":false},{\"title\":\"课程组件\",\"showTitle\":2,\"height\":\"50\",\"eduIds\":\"2,3,4,5\",\"style\":4,\"isOpen\":false}]', '2021-04-12 20:29:03', '2021-12-20 22:57:38', '首页设计');
INSERT INTO `sys_config` VALUES ('23', '积分设置', 'IntegralSetting', '{\"registerIntegral\":300,\"signInIntegral\":\"500-5000\",\"contentIntegral\":50,\"contentCount\":2,\"groupIntegral\":50,\"groupCount\":2,\"answerIntegral\":50,\"answerCount\":2,\"commentIntegral\":50,\"commentCount\":2,\"likefavoriteIntegral\":50,\"likefavoriteCount\":2,\"followIntegral\":50,\"followCount\":2,\"reportIntegral\":50,\"reportCount\":2}', '2021-04-22 20:57:25', '2021-09-20 16:12:14', '积分设置');
INSERT INTO `sys_config` VALUES ('24', '通知设置', 'NoticeSetting', '{\"register\":\"{siteTitle}欢迎您的加入\",\"create\":\"您发布《{title}》,已经审核: {reason}\",\"remove\":\"您发布的内容{title}，被删除了，{reason}\",\"groupCreate\":\"您创建《{title}》,已经审核: {reason}\",\"groupRemove\":\"您创建{title}，被删除了: {reason}\",\"report\":\"您举报的内容，已处理，{reason}\",\"userProhibit\":\"{reason}\",\"verify\":\"您实名认证已审核，{reason}\",\"cash\":\"您编号为{code}提现申请已经打款，\"}', '2021-05-16 02:07:09', '2021-06-21 00:44:05', '通知设置');
INSERT INTO `sys_config` VALUES ('26', '登录注册设置', 'AuthSetting', '{\"registerMode\":\"email\",\"policyUrl\":\"\",\"protocolUrl\":\"\",\"social\":[],\"register\":\"恭喜您注册成为我们的用户\"}', '2021-06-22 00:51:45', '2022-01-06 14:59:51', '登录注册设置');
INSERT INTO `sys_config` VALUES ('27', '阿里云oss配置', 'AlyOssOption', '{\"endpoint\":\"fibercms.oss-cn-guangzhou.aliyuncs.com\",\"accessKeyId\":\"LTAI5tFGYMtpuMG8xiDUxyQa\",\"accessKeySecret\":\"CIH3VRjTj9SIXgTRn0I7P8onykqUex\",\"bucketName\":\"fibercms\"}', '2021-07-02 21:38:09', '2021-12-18 14:07:09', '阿里云oss配置');
INSERT INTO `sys_config` VALUES ('28', '七牛oss配置', 'QiNiuOssOption', '{\"endpoint\":\"21312\",\"accessKeyId\":\"312312\",\"accessKeySecret\":\"3123\",\"bucketName\":\"123\",\"address\":\"213123\"}', '2021-07-02 22:17:48', '2021-07-02 22:17:48', '七牛oss配置');
INSERT INTO `sys_config` VALUES ('29', '阿里云短信配置', 'AlySmsOptions', '{\"id\":\"LTAIyOnBE3wGtyAA\",\"secret\":\"PlwvkrExsup3XccmMxaonmJe4HshK3\",\"publicKey\":null,\"appPublicKey\":null,\"alyCertPublicKey\":null,\"alyRootCert\":null}', '2021-11-28 17:05:20', '2021-11-28 17:05:27', '阿里云短信配置');
INSERT INTO `sys_config` VALUES ('30', 'H5首页设计', 'H5Design', '[{\"title\":\"轮播1\",\"height\":360,\"list\":[{\"isPlatform\":1,\"module\":\"article\",\"id\":\"1\",\"link\":\"https://www.baidu.com/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg\"},{\"link\":\"/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftspud91rg2vgrxnm.jpg\",\"isPlatform\":1,\"module\":\"video\",\"id\":\"2\"},{\"link\":\"/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftrximm0wdbicgjmj.jpg\",\"isPlatform\":1,\"module\":\"audio\",\"id\":\"3\"},{\"link\":\"/\",\"cover\":\"http://localhost:8199/public/uploads/2021-11-19/cftsj83cepnpomdgky.jpg\",\"isPlatform\":1,\"module\":\"resource\",\"id\":\"3\"}],\"style\":1},{\"title\":\"社区组件\",\"showTitle\":2,\"topicIds\":\"9,8,7\",\"questionIds\":\"\",\"style\":2},{\"title\":\"自定义内容组件\",\"showTitle\":2,\"list\":[{\"title\":\"测试\",\"icon\":\"\",\"link\":\"http://fibercms.com/\",\"isPlatform\":1},{\"title\":\"测试\",\"icon\":\"\",\"link\":\"http://fibercms.com/\",\"isPlatform\":1},{\"title\":\"测试\",\"icon\":\"\",\"link\":\"http://fibercms.com/\",\"isPlatform\":1},{\"title\":\"测试\",\"icon\":\"\",\"link\":\"http://fibercms.com/\",\"isPlatform\":1}],\"style\":5},{\"title\":\"课程组件\",\"showTitle\":2,\"eduIds\":\"2,3,4\",\"style\":4},{\"title\":\"自定义图片组件\",\"showTitle\":2,\"image\":\"http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg\",\"link\":\"http://fibercms.com/\",\"style\":6,\"isPlatform\":1},{\"title\":\"投稿组件\",\"showTitle\":2,\"audioIds\":\"2,3,4\",\"videoIds\":\"2,3,4\",\"resourceIds\":\"2,3,4\",\"style\":7},{\"title\":\"文章组件\",\"showTitle\":2,\"articleIds\":\"2,3,4,1\",\"style\":3}]', '2021-11-30 12:39:22', '2021-11-30 20:01:24', 'H5首页设计');
INSERT INTO `sys_config` VALUES ('31', '微信小程序首页设计', 'WechatMiniProgramDesign', '[{\"title\":\"轮播1\",\"height\":360,\"list\":[{\"isPlatform\":2,\"module\":\"article\",\"id\":\"1\"},{\"link\":\"\",\"cover\":\"\",\"isPlatform\":2,\"module\":\"video\",\"id\":\"2\"},{\"link\":\"\",\"cover\":\"\",\"isPlatform\":2,\"module\":\"audio\",\"id\":\"3\"},{\"link\":\"\",\"cover\":\"\",\"isPlatform\":2,\"module\":\"resource\",\"id\":\"3\"}],\"style\":1},{\"title\":\"社区组件\",\"showTitle\":2,\"topicIds\":\"9,8,7\",\"questionIds\":\"\",\"style\":2},{\"title\":\"自定义内容组件\",\"showTitle\":2,\"list\":[{\"title\":\"动态\",\"icon\":\"\",\"link\":\"pages/feed/index\",\"isPlatform\":1},{\"title\":\"动态\",\"icon\":\"\",\"link\":\"pages/feed/index\",\"isPlatform\":1},{\"title\":\"动态\",\"icon\":\"\",\"link\":\"pages/feed/index\",\"isPlatform\":1},{\"title\":\"动态\",\"icon\":\"\",\"link\":\"pages/feed/index\",\"isPlatform\":1}],\"style\":5},{\"title\":\"课程组件\",\"showTitle\":2,\"eduIds\":\"2,3,4\",\"style\":4},{\"title\":\"自定义图片组件\",\"showTitle\":2,\"image\":\"http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg\",\"link\":\"http://fibercms.com/\",\"style\":6},{\"title\":\"投稿组件\",\"showTitle\":2,\"audioIds\":\"2,3,4\",\"videoIds\":\"2,3,4\",\"resourceIds\":\"2,3,4\",\"style\":7},{\"title\":\"文章组件\",\"showTitle\":2,\"articleIds\":\"2,3,4,1\",\"style\":3}]', '2021-11-30 14:34:02', '2021-11-30 14:44:09', '微信小程序首页设计');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `create_by` bigint(20) DEFAULT NULL,
  `dict_sort` int(4) DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(4) DEFAULT '2' COMMENT '是否默认 1是 2否',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES ('1', '1', '1', '男', '0', 'sys_user_sex', '', '', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '性别男');
INSERT INTO `sys_dict_data` VALUES ('2', '1', '2', '女', '1', 'sys_user_sex', '', '', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '性别女');
INSERT INTO `sys_dict_data` VALUES ('3', '1', '3', '未知', '2', 'sys_user_sex', '', '', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '性别未知');
INSERT INTO `sys_dict_data` VALUES ('4', '1', '1', '显示', '0', 'sys_show_hide', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '显示菜单');
INSERT INTO `sys_dict_data` VALUES ('5', '1', '2', '隐藏', '1', 'sys_show_hide', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '隐藏菜单');
INSERT INTO `sys_dict_data` VALUES ('6', '1', '1', '正常', '0', 'sys_normal_disable', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES ('7', '1', '2', '停用', '1', 'sys_normal_disable', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES ('8', '1', '1', '正常', '0', 'sys_job_status', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES ('9', '1', '2', '暂停', '1', 'sys_job_status', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES ('10', '1', '1', '默认', 'DEFAULT', 'sys_job_group', '', '', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '默认分组');
INSERT INTO `sys_dict_data` VALUES ('11', '1', '2', '系统', 'SYSTEM', 'sys_job_group', '', '', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '系统分组');
INSERT INTO `sys_dict_data` VALUES ('12', '1', '1', '是', 'Y', 'sys_yes_no', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '系统默认是');
INSERT INTO `sys_dict_data` VALUES ('13', '1', '2', '否', 'N', 'sys_yes_no', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '系统默认否');
INSERT INTO `sys_dict_data` VALUES ('14', '1', '1', '通知', '1', 'sys_notice_type', '', 'warning', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '通知');
INSERT INTO `sys_dict_data` VALUES ('15', '1', '2', '公告', '2', 'sys_notice_type', '', 'success', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '公告');
INSERT INTO `sys_dict_data` VALUES ('16', '1', '1', '正常', '0', 'sys_notice_status', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES ('17', '1', '2', '关闭', '1', 'sys_notice_status', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '关闭状态');
INSERT INTO `sys_dict_data` VALUES ('18', '1', '1', '新增', '1', 'sys_oper_type', '', 'info', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '新增操作');
INSERT INTO `sys_dict_data` VALUES ('19', '1', '2', '修改', '2', 'sys_oper_type', '', 'info', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '修改操作');
INSERT INTO `sys_dict_data` VALUES ('20', '1', '3', '删除', '3', 'sys_oper_type', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '删除操作');
INSERT INTO `sys_dict_data` VALUES ('21', '1', '4', '授权', '4', 'sys_oper_type', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '授权操作');
INSERT INTO `sys_dict_data` VALUES ('22', '1', '5', '导出', '5', 'sys_oper_type', '', 'warning', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '导出操作');
INSERT INTO `sys_dict_data` VALUES ('23', '1', '6', '导入', '6', 'sys_oper_type', '', 'warning', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '导入操作');
INSERT INTO `sys_dict_data` VALUES ('24', '1', '7', '强退', '7', 'sys_oper_type', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '强退操作');
INSERT INTO `sys_dict_data` VALUES ('25', '1', '8', '生成代码', '8', 'sys_oper_type', '', 'warning', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '生成操作');
INSERT INTO `sys_dict_data` VALUES ('26', '1', '9', '清空数据', '9', 'sys_oper_type', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '清空操作');
INSERT INTO `sys_dict_data` VALUES ('27', '1', '1', '成功', '0', 'sys_common_status', '', 'primary', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES ('28', '1', '2', '失败', '1', 'sys_common_status', '', 'danger', null, '0', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES ('31', '1', '0', '阴阳人', '6', 'sys_user_sex', '', '', null, '0', '2020-03-20 19:49:06', null, '性别阴阳人');
INSERT INTO `sys_dict_data` VALUES ('33', '1', '0', '本地', '1', 'sys_media_engine', '', '', null, '0', '2020-04-10 19:06:55', null, '存储引擎本地');
INSERT INTO `sys_dict_data` VALUES ('34', '1', '0', '阿里云', '2', 'sys_media_engine', '', '', null, '0', '2020-04-10 19:07:10', null, '阿里云OSS');
INSERT INTO `sys_dict_data` VALUES ('35', '1', '0', '七牛云', '3', 'sys_media_engine', '', '', null, '0', '2020-04-10 19:07:31', null, '七牛云oss');
INSERT INTO `sys_dict_data` VALUES ('43', '1', '0', '圈子', 'circle', 'sys_modules', '', '', null, '0', '2020-04-10 21:39:57', null, '圈子模块');
INSERT INTO `sys_dict_data` VALUES ('44', '1', '0', '资源', 'resource', 'sys_modules', '', '', null, '0', '2020-04-10 21:41:54', '2020-04-19 23:01:01', '资源模块');
INSERT INTO `sys_dict_data` VALUES ('45', '1', '0', '帖子', 'feed', 'sys_modules', '', '', null, '0', '2020-04-10 21:42:17', null, '帖子模块');
INSERT INTO `sys_dict_data` VALUES ('61', '1', '0', '编曲', 'bianqu', 'sys_verify_skill', '', '', null, '0', '2020-06-03 01:32:26', null, '');
INSERT INTO `sys_dict_data` VALUES ('62', '1', '0', '混音', 'hunyin', 'sys_verify_skill', '', '', null, '0', '2020-06-03 01:32:42', null, '');
INSERT INTO `sys_dict_data` VALUES ('63', '1', '0', '设计', 'sheji', 'sys_verify_skill', '', '', null, '0', '2020-06-03 01:33:06', null, '');
INSERT INTO `sys_dict_data` VALUES ('64', '1', '0', '后期', 'houqi', 'sys_verify_skill', '', '', null, '0', '2020-06-03 01:33:38', null, '');
INSERT INTO `sys_dict_data` VALUES ('65', '1', '0', '开启', '0', 'sys_user_verify', '', '', null, '0', '2020-10-24 23:24:47', null, '');
INSERT INTO `sys_dict_data` VALUES ('66', '1', '0', '禁用', '1', 'sys_user_verify', '', '', null, '0', '2020-10-24 23:25:04', null, '');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `title` varchar(100) DEFAULT '' COMMENT '字典名称',
  `type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态（0正常，2停用）',
  `create_by` bigint(20) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE KEY `dict_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES ('1', '用户性别', 'sys_user_sex', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '用户性别列表');
INSERT INTO `sys_dict_type` VALUES ('2', '菜单状态', 'sys_show_hide', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '菜单状态列表');
INSERT INTO `sys_dict_type` VALUES ('3', '系统开关', 'sys_normal_disable', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '系统开关列表');
INSERT INTO `sys_dict_type` VALUES ('4', '任务状态', 'sys_job_status', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '任务状态列表');
INSERT INTO `sys_dict_type` VALUES ('5', '任务分组', 'sys_job_group', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '任务分组列表');
INSERT INTO `sys_dict_type` VALUES ('6', '系统是否', 'sys_yes_no', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '系统是否列表');
INSERT INTO `sys_dict_type` VALUES ('7', '通知类型', 'sys_notice_type', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '通知类型列表');
INSERT INTO `sys_dict_type` VALUES ('8', '通知状态', 'sys_notice_status', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '通知状态列表');
INSERT INTO `sys_dict_type` VALUES ('9', '操作类型', 'sys_oper_type', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '操作类型列表');
INSERT INTO `sys_dict_type` VALUES ('10', '系统状态', 'sys_common_status', '0', '1', '2018-03-16 11:33:00', '2018-03-16 11:33:00', '登录状态列表');
INSERT INTO `sys_dict_type` VALUES ('15', '文件引擎', 'sys_media_engine', '0', '1', '2020-04-10 19:04:52', '2020-04-10 19:11:04', '文件引擎列表');
INSERT INTO `sys_dict_type` VALUES ('16', '文件类型', 'sys_media_type', '0', '1', '2020-04-10 19:12:03', null, '文件类型列表');
INSERT INTO `sys_dict_type` VALUES ('17', '系统模块', 'sys_modules', '0', '1', '2020-04-10 21:38:02', '2020-04-14 16:50:49', '系统模块列表');
INSERT INTO `sys_dict_type` VALUES ('24', '职业技能', 'sys_verify_skill', '0', '1', '2020-06-03 01:30:26', null, '认证用户所用的职业技能');
INSERT INTO `sys_dict_type` VALUES ('25', '认证服务', 'sys_user_verify', '0', '1', '2020-10-24 23:23:47', null, '实名认证列表');

-- ----------------------------
-- Table structure for sys_edu
-- ----------------------------
DROP TABLE IF EXISTS `sys_edu`;
CREATE TABLE `sys_edu` (
  `edu_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `cate_id` bigint(20) DEFAULT NULL COMMENT '分类',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `content` text,
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `section` text COMMENT '视频地址',
  `max` int(10) DEFAULT NULL,
  `favorites` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `joins` bigint(20) DEFAULT NULL,
  `hots` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL COMMENT '播放量',
  `type` tinyint(4) DEFAULT NULL COMMENT '课程类型 1线下，2线上',
  `join_mode` tinyint(4) DEFAULT '1' COMMENT '查看权限 1公开下载，2付费下载',
  `price` decimal(10,2) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `delete_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`edu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_edu
-- ----------------------------
INSERT INTO `sys_edu` VALUES ('2', '1', '7', 'Photoshop 平面设计完整基础学习课程视频教程（英文）Monster Level In Photoshop – The best tutorial for Beginners', 'http://localhost:8199/public/uploads/2021-11-19/cftrr6s1ecekzzmlt6.jpg', '<h2>介绍</h2>\n<p>MP4 |视频：h264，1280x720 |音频: AAC, 44.1 KHz, 2 Ch<br />类型：在线学习 |语言：英语+srt |持续时间：24 个讲座 (5h 29m) |大小：3.54 GB</p>\n<p>Photoshop的完整解释</p>\n<p>你会学到什么：<br />您将学习所有最重要的工具，以及如何编辑照片、修饰、裁剪社交媒体图片、更改背景等等。<br />完整而专业地解释程序设置。<br />您将知道如何润饰照片以去除瑕疵、修复红眼、美白牙齿等。<br />您将学习如何保存 Photoshop 项目以供打印和 Web 使用。<br />您将学习如何使用图层面板，包括创建和编辑图层蒙版。<br />您将学习如何使用多种选择工具选择和编辑图像的一部分。<br />学习有用的键盘快捷键和最佳实践<br />您将创造出惊人的逼真阴影。</p>\n<p>要求<br />不需要Photoshop的先前知识。<br />任何版本的 Adob​​e Photoshop，最好不早于 Photoshop CS6。</p>\n<p>描述<br />大家好，我叫 Ismail Moqbel。在您学习 Photoshop 程序的整个过程中，我将成为您的老师。</p>\n<p>我想指出的是，我的母语不是英语，我曾尝试用清晰简单的语言讲授这门课程，但我确信我犯了几个语法和发音错误，所以请原谅我。</p>\n<p>您是否正在为自学 Photoshop 而苦苦挣扎和遇到困难？本课程将使您能够专业地使用 Photoshop。</p>\n<p>在本课程中，我将教您有关开始使用 Photoshop 的知识。您将学习如何使用 Photoshop 进行平面设计和 Photoshop 润饰。本课程将帮助您获得理想的工作。</p>\n<p>如果您以前从未打开过 Photoshop，或者您已经打开过 Photoshop 并且在基础知识方面苦苦挣扎，请跟随我，我们将一起学习如何使用 Photoshop 制作精美的图像。</p>\n<p>本课程适用于初学者和中级。您不需要任何 Photoshop、摄影或设计方面的先验知识。我们将逐步开始工作。您将神奇地增强我们的背景，并在必要时从图像中完全删除人物。</p>\n<p>这门课程的结构非常好 - 从简单到难的主题。也有很多详细的解释。</p>\n<p>本课程适用于：</p>\n<p>本课程将是任何想要在职业生涯中开始使用 Photoshop 并通过 Photoshop 技能获得报酬的人以及任何想要设计自己的图形和从头开始编辑自己的照片的创意人员的理想课程</p>\n<p>它将更新到 2022 年</p>\n<p>本课程适用于谁<br />初学者<br />中间的<br />13岁以上</p>\n<p>MP4 | Video: h264, 1280x720 | Audio: AAC, 44.1 KHz, 2 Ch<br />Genre: eLearning | Language: English + srt | Duration: 24 lectures (5h 29m) | Size: 3.54 GB</p>\n<p>Full explanation of Photoshop</p>\n<p>What you\'ll learn:<br />You\'ll learn all of the most important tools, plus how to edit photos, do touch-ups, crop social media pics, change backgrounds and more.<br />Explaining the program settings completely and professionally.<br />You\'ll know how to retouch photos to remove blemishes, fix red-eye, whiten teeth, and more.<br />You\'ll learn how to save your Photoshop projects for print and web.<br />You\'ll learn how to use the layers panel including creating and editing layer masks.<br />You\'ll learn how to select and edit just parts of your image with a number of selection tools.<br />Learn useful keyboard shortcuts and best practices<br />You will create amazingly realistic shadows.</p>\n<p>Requirements<br />No previous knowledge of Photoshop is required.<br />Any version of Adobe Photoshop, preferably not older than Photoshop CS6.</p>\n<p>Description<br />Hi there, my name is Ismail Moqbel. I will be your teacher throughout your journey to learn Photoshop program.</p>\n<p>I would like to point that I am not native English speaker , I have tried delivering this course in a clear and simple language but I am sure that I made several grammatical and pronunciation mistakes , so please do forgive me for that.</p>\n<p>Are you struggling and encounter difficulties to learn Photoshop on your own? This course will allow you to use Photoshop professionally.</p>\n<p>In this course I will teach you what you need to know about getting started with Photoshop. You\'ll learn how to use Photoshop for use in Graphic Design &amp; for Photoshop Retouching. This course will help your get your dream job.</p>\n<p>If you have never opened Photoshop before or you have already opened Photoshop and are struggling with the basics, follow me and together we will learn how to make beautiful images using Photoshop.</p>\n<p>This course is for beginners and intermediate. You do not need any prior knowledge of Photoshop, photography or design. We will start work our way through step by step. You will magically enhance our background and when necessary completely remove people from images.</p>\n<p>This course is very good structured - from easy to hard topics. Also it has a lot of detailed explanations .</p>\n<p>Who this course is for:</p>\n<p>This course will be an ideal course for anyone who wants to start using Photoshop in their career &amp; get paid for their Photoshop skills and any creatives who want to design their own graphics and edit their own photos from scratch</p>\n<p>It will be updated up to a year 2022</p>\n<p>Who this course is for<br />beginners<br />intermediate<br />Over 13 years old</p>', '测试课程2', '[{\"children\":[{\"isWatch\":false,\"link\":\"http://localhost:8199/public/uploads/2021-08-26/cdtl3q3yml0ezcgnhs.mp4\",\"title\":\"第2节课\"},{\"isWatch\":true,\"link\":\"http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4\",\"title\":\"第1课时\"}],\"title\":\"第1章节\"}]', '35435', '123123', '123', '1121', '123', '1237', '2', '2', '213.00', '2', null, '2021-07-12 14:15:36', '2021-11-19 20:56:42', '');
INSERT INTO `sys_edu` VALUES ('3', '1', '7', '浅谈角色设计思路', 'http://localhost:8199/public/uploads/2021-11-19/cftrximm0wdbicgjmj.jpg', '<p>测试课程快递发给你是打设计的小思路有很多，目的其实就是给我们提供一个设计的切入点。给大家分享一个我自己总结的让设计比较生动的思路。A+B+C的思路：&nbsp; A.人物的第一性征，性别，年龄，形体，性格，基本的职业 B.人物的第二性征，尝试和其他元素的结合，比如常用的有异变的元素，加入朋克或者机械的元素，或者是和一些其他的职业特征和时髦元素进行混搭。 C.人物的第三性征，故事性，就是这个角色经历了什么我们能否尝试在角色本身体现出来，比如脸上的疤痕，被火焰熏黑的武器，湿透的下半身。</p>', '过来看看\n', '[{\"children\":[{\"isWatch\":false,\"link\":\"http://localhost:8199/public/uploads/2021-08-26/cdtl3q3yml0ezcgnhs.mp4\",\"title\":\"第2节课\"},{\"isWatch\":true,\"link\":\"http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4\",\"title\":\"第1课时\"}],\"title\":\"第1章节\"}]', '1000', '3123', '12312', '13123', '231', '3218', '2', '1', '231.00', '2', null, '2021-08-25 01:03:40', '2021-11-19 21:03:49', '');
INSERT INTO `sys_edu` VALUES ('4', '1', '7', 'Kontakt 6.2.2安装视频教程-专治各种入库和相关问题', 'http://localhost:8199/public/uploads/2021-11-19/cfts2dgbpc7lqbrs9o.jpg', '<p>小编点评：Kontakt 6的相关安装视频教程！因为录制得伧促，不好之处请见谅！有问题请留言！我会第一时间解答的！如有需要出使用教程的，我会看需求的人数来定吧！</p>', 'Kontakt 6的相关安装视频教程！因为录制得伧促，不好之处请见谅！有问题请留言！我会第一时间解答的！如有需要出使用教程的', '[{\"children\":[{\"isWatch\":false,\"link\":null,\"title\":\"第1课时\"}],\"title\":\"第1章节\"}]', '999', '1231', '231', '312', '12123', '2323', '2', '1', '0.00', '2', null, '2021-11-19 21:07:54', '2021-11-19 21:07:54', '');
INSERT INTO `sys_edu` VALUES ('5', '1', '7', '母带神器臭氧9自动母带iZotope Ozone Advanced 9安装教程', 'http://localhost:8199/public/uploads/2021-11-19/cfts35arimgovyddvw.png', '<p>风格化风格化是豆腐干豆腐干士大夫敢死队</p>', 'hghjbgjklhlkjhjknikmljiuj', '[{\"children\":[{\"isWatch\":false,\"link\":null,\"title\":\"第1课时\"}],\"title\":\"第1章节\"}]', '435345', '123', '31231', '123', '21312', '23147', '1', '1', '0.00', '2', null, '2021-11-19 21:08:47', '2021-11-19 21:16:16', '');

-- ----------------------------
-- Table structure for sys_grade
-- ----------------------------
DROP TABLE IF EXISTS `sys_grade`;
CREATE TABLE `sys_grade` (
  `grade_id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(30) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `integral` int(10) DEFAULT NULL,
  `create_group` int(10) DEFAULT NULL,
  `posts_module` varchar(255) DEFAULT NULL,
  `common_auth` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`grade_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_grade
-- ----------------------------
INSERT INTO `sys_grade` VALUES ('1', '等级一', 'http://localhost:8199/public/uploads/2021-07-22/cczp8yet9cxwo178kk.png', '20', '5', '[\"audio\",\"article\",\"video\",\"resource\",\"edu\"]', '[\"answer\",\"report\",\"comment\",\"group\",\"question\",\"topic\",\"upload\"]');
INSERT INTO `sys_grade` VALUES ('2', '等级二', 'http://localhost:8199/public/uploads/2021-07-22/cczp8yet9cxwo178kk.png', '231', '123', '[\"video\",\"article\"]', '[\"comment\",\"answer\",\"report\"]');
INSERT INTO `sys_grade` VALUES ('3', '等级三', 'http://localhost:8199/public/uploads/2021-07-22/cczp8yet9cxwo178kk.png', '545', '5465', '[\"article\",\"video\",\"resource\",\"edu\"]', '[\"comment\",\"answer\",\"upload\",\"topic\"]');

-- ----------------------------
-- Table structure for sys_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_group`;
CREATE TABLE `sys_group` (
  `group_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '创建人',
  `cate_id` bigint(20) DEFAULT NULL COMMENT '小组分类',
  `title` varchar(50) DEFAULT NULL COMMENT '小组名称',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `join_mode` tinyint(4) DEFAULT '1' COMMENT '小组类型1 公共小组，2付费小组，3专属小组',
  `price` decimal(10,2) DEFAULT NULL COMMENT '费用',
  `secret_key` varchar(100) DEFAULT NULL COMMENT '加入角色',
  `icon` varchar(255) DEFAULT NULL COMMENT '小组图标',
  `joins` bigint(20) DEFAULT NULL,
  `hots` bigint(20) DEFAULT NULL,
  `contents` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL COMMENT '小组描述',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态 0 全部  1待审   2通过  3拒绝',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '跟新时间',
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_group
-- ----------------------------
INSERT INTO `sys_group` VALUES ('2', '3', '9', '意见反馈', 'http://localhost:8199/public/uploads/2021-08-26/cdtlbxdhn37reyybb4.jpg', '1', '0.00', 'a121300', 'http://localhost:8199/public/uploads/2021-08-26/cdtlbxdhn37reyybb4.jpg', '10', '21312', '10', '174', 'bug反馈，功能意见', '', '2', '2021-07-12 21:45:06', '2021-08-26 18:38:54', null);
INSERT INTO `sys_group` VALUES ('3', '1', '9', '招聘', 'http://localhost:8199/public/uploads/2021-08-26/cdtlbxdhn37reyybb4.jpg', '1', '0.00', 'a121300', '26/cdtlbxdhn37reyybb4.jpg', '12312', '12312', '12', '1276', 'bug反馈，功能意见', null, '2', '2021-12-04 16:14:28', '2021-12-04 16:14:33', null);

-- ----------------------------
-- Table structure for sys_manger
-- ----------------------------
DROP TABLE IF EXISTS `sys_manger`;
CREATE TABLE `sys_manger` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nick_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(4) DEFAULT '3' COMMENT '用户性别（1男 2女 3未知）',
  `avatar` varchar(255) DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `salt` char(10) DEFAULT NULL COMMENT '密码盐',
  `status` tinyint(4) DEFAULT '0' COMMENT '帐号状态（1停用,2正常）',
  `login_ip` varchar(50) DEFAULT '' COMMENT '最后登陆IP',
  `login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `delete_time` datetime DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_manger
-- ----------------------------
INSERT INTO `sys_manger` VALUES ('1', '我最大', 'fiber@admin.com', '', '1', 'http://localhost:8199/public/uploads/2021-08-26/cdtkto2n81q4lcaqgx.jpg', 'f3565ae9276f3b7244230c1bcbb4ff4c', 'ePJsln', '2', '[::1]', '2022-01-27 13:58:47', '2021-05-10 05:50:02', '2021-08-26 18:39:21', '', null);
INSERT INTO `sys_manger` VALUES ('3', '新用户248752', 'test@admin.com', '', '1', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', '528fae61d5899c4eb7edc13bc9ed1945', 'Zh0DRf', '2', '[::1]', '2022-01-27 13:58:28', '2021-05-10 05:58:47', '2021-11-20 00:19:07', '', null);
INSERT INTO `sys_manger` VALUES ('4', '新用户869377', 'mushokumunou@gmail.com', '', '3', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', '75965be72c6c21f7345789ce547c2ec9', 'gCiJ6S', '2', '178.132.6.37', '2021-05-11 10:40:24', '2021-05-11 10:40:14', '2021-09-28 13:34:10', '', null);

-- ----------------------------
-- Table structure for sys_manger_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_manger_role`;
CREATE TABLE `sys_manger_role` (
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_manger_role
-- ----------------------------
INSERT INTO `sys_manger_role` VALUES ('1', '1');
INSERT INTO `sys_manger_role` VALUES ('3', '15');
INSERT INTO `sys_manger_role` VALUES ('4', '1');

-- ----------------------------
-- Table structure for sys_media
-- ----------------------------
DROP TABLE IF EXISTS `sys_media`;
CREATE TABLE `sys_media` (
  `media_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '上传的用户',
  `link` varchar(255) DEFAULT NULL COMMENT '文件链接',
  `path` varchar(255) DEFAULT NULL COMMENT '存放路径',
  `name` varchar(50) DEFAULT NULL COMMENT '文件名称',
  `or_name` varchar(50) DEFAULT NULL COMMENT '原始文件名称',
  `size` varchar(50) NOT NULL COMMENT '文件大小',
  `upload_key` tinyint(4) NOT NULL COMMENT '上传方式 1 为本地上传， 2为oss上传',
  `ext` varchar(10) DEFAULT NULL COMMENT '文件后缀',
  `status` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `delete_time` datetime DEFAULT NULL,
  `media_type` varchar(20) DEFAULT NULL COMMENT '文件类型',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`media_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=283 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_media
-- ----------------------------
INSERT INTO `sys_media` VALUES ('163', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftkh4fqsdcgux3vta.png', '/public/uploads/2021-11-19/cftkh4fqsdcgux3vta.png', 'cftkh4fqsdcgux3vta.png', '默认文件1637305347696.png', '38134', '1', '.png', '2', '2021-11-19 15:10:25', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('164', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftlm1x82z1ogbe0bm.png', '/public/uploads/2021-11-19/cftlm1x82z1ogbe0bm.png', 'cftlm1x82z1ogbe0bm.png', 'photoshop.png', '3046', '1', '.png', '2', '2021-11-19 16:03:52', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('167', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftn886w6pso93otxw.png', '/public/uploads/2021-11-19/cftn886w6pso93otxw.png', 'cftn886w6pso93otxw.png', 'snipaste.png', '10379', '1', '.png', '2', '2021-11-19 17:19:51', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('168', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftn8mjwg6t9iaomlf.png', '/public/uploads/2021-11-19/cftn8mjwg6t9iaomlf.png', 'cftn8mjwg6t9iaomlf.png', 'procreate.png', '8550', '1', '.png', '2', '2021-11-19 17:20:22', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('169', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnb42nzy21h5aw49.png', '/public/uploads/2021-11-19/cftnb42nzy21h5aw49.png', 'cftnb42nzy21h5aw49.png', 'krita.png', '14866', '1', '.png', '2', '2021-11-19 17:23:37', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('170', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnc1l4tz1zml3d8h.png', '/public/uploads/2021-11-19/cftnc1l4tz1zml3d8h.png', 'cftnc1l4tz1zml3d8h.png', 'artstudio.png', '7691', '1', '.png', '2', '2021-11-19 17:24:50', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('171', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftncplplm0law4plg.png', '/public/uploads/2021-11-19/cftncplplm0law4plg.png', 'cftncplplm0law4plg.png', 'live2d.png', '8974', '1', '.png', '2', '2021-11-19 17:25:43', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('172', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnonwtwzmjcc0hwi.jpg', '/public/uploads/2021-11-19/cftnonwtwzmjcc0hwi.jpg', 'cftnonwtwzmjcc0hwi.jpg', '20211010-JsiZFi.jpg', '18953', '1', '.jpg', '2', '2021-11-19 17:41:19', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('173', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', '/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', 'cftnoto7hqia1j1rct.mp4', '8831bfcb00c17546e_1_circle.mp4', '956810', '1', '.mp4', '2', '2021-11-19 17:41:32', null, 'MP4', '');
INSERT INTO `sys_media` VALUES ('174', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnpvc6xash3icb7t.jpg', '/public/uploads/2021-11-19/cftnpvc6xash3icb7t.jpg', 'cftnpvc6xash3icb7t.jpg', '20211011-OARliX.jpg', '20999', '1', '.jpg', '2', '2021-11-19 17:42:54', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('175', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnqt0817ekqffczc.jpg', '/public/uploads/2021-11-19/cftnqt0817ekqffczc.jpg', 'cftnqt0817ekqffczc.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '2', '2021-11-19 17:44:07', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('176', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnsb3a88t58v3axf.png', '/public/uploads/2021-11-19/cftnsb3a88t58v3axf.png', 'cftnsb3a88t58v3axf.png', 'QQ图片20211119174548.png', '173737', '1', '.png', '2', '2021-11-19 17:46:05', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('177', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnvggid63b6lts3p.png', '/public/uploads/2021-11-19/cftnvggid63b6lts3p.png', 'cftnvggid63b6lts3p.png', 'QQ图片20211119174948.png', '295376', '1', '.png', '2', '2021-11-19 17:50:12', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('178', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', '/public/uploads/2021-11-19/cftnwa2sj2upfuoapv.mp3', 'cftnwa2sj2upfuoapv.mp3', '8477751_1623867216.mp3', '3683175', '1', '.mp3', '2', '2021-11-19 17:51:16', null, 'MP3', '');
INSERT INTO `sys_media` VALUES ('179', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftny9chbl1amnxtzl.png', '/public/uploads/2021-11-19/cftny9chbl1amnxtzl.png', 'cftny9chbl1amnxtzl.png', 'QQ图片20211119175334.png', '329232', '1', '.png', '2', '2021-11-19 17:53:51', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('181', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfto04fi4z4hqz0iea.png', '/public/uploads/2021-11-19/cfto04fi4z4hqz0iea.png', 'cfto04fi4z4hqz0iea.png', 'QQ图片20211119175610.png', '333367', '1', '.png', '2', '2021-11-19 17:56:17', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('182', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfto13hrsbm66lshhl.png', '/public/uploads/2021-11-19/cfto13hrsbm66lshhl.png', 'cfto13hrsbm66lshhl.png', 'QQ图片20211119175720.png', '113120', '1', '.png', '2', '2021-11-19 17:57:34', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('183', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfto1rq6w4rym2ov9t.png', '/public/uploads/2021-11-19/cfto1rq6w4rym2ov9t.png', 'cfto1rq6w4rym2ov9t.png', 'QQ图片20211119175814.png', '230559', '1', '.png', '2', '2021-11-19 17:58:26', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('184', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfto2qrr3d25zm0lr9.png', '/public/uploads/2021-11-19/cfto2qrr3d25zm0lr9.png', 'cfto2qrr3d25zm0lr9.png', 'QQ图片20211119175936.png', '329764', '1', '.png', '2', '2021-11-19 17:59:43', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('185', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftrr6s1ecekzzmlt6.jpg', '/public/uploads/2021-11-19/cftrr6s1ecekzzmlt6.jpg', 'cftrr6s1ecekzzmlt6.jpg', '20211026-NOnYqA.jpg', '10657', '1', '.jpg', '2', '2021-11-19 20:52:42', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('186', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftrximm0wdbicgjmj.jpg', '/public/uploads/2021-11-19/cftrximm0wdbicgjmj.jpg', 'cftrximm0wdbicgjmj.jpg', '1626943999-boyu.jpg', '83681', '1', '.jpg', '2', '2021-11-19 21:00:57', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('187', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfts2dgbpc7lqbrs9o.jpg', '/public/uploads/2021-11-19/cfts2dgbpc7lqbrs9o.jpg', 'cfts2dgbpc7lqbrs9o.jpg', '8285710645fd9ca7.jpg', '81341', '1', '.jpg', '2', '2021-11-19 21:07:18', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('188', '1', 'http://localhost:8199/public/uploads/2021-11-19/cfts35arimgovyddvw.png', '/public/uploads/2021-11-19/cfts35arimgovyddvw.png', 'cfts35arimgovyddvw.png', 'QQ图片20211119210811.png', '319537', '1', '.png', '2', '2021-11-19 21:08:19', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('189', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftsj83cepnpomdgky.jpg', '/public/uploads/2021-11-19/cftsj83cepnpomdgky.jpg', 'cftsj83cepnpomdgky.jpg', 'samsungqdoled_2_large.jpg', '239267', '1', '.jpg', '2', '2021-11-19 21:29:19', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('190', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftsjqv0421pcqd3j1.jpg', '/public/uploads/2021-11-19/cftsjqv0421pcqd3j1.jpg', 'cftsjqv0421pcqd3j1.jpg', 'psremoteplay2021.jpg', '74147', '1', '.jpg', '2', '2021-11-19 21:29:59', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('191', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftskpi0mwiskcldfg.jpg', '/public/uploads/2021-11-19/cftskpi0mwiskcldfg.jpg', 'cftskpi0mwiskcldfg.jpg', 'bdcdae5bc1464515891f34.jpg', '927334', '1', '.jpg', '2', '2021-11-19 21:31:15', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('192', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftspud91rg2vgrxnm.jpg', '/public/uploads/2021-11-19/cftspud91rg2vgrxnm.jpg', 'cftspud91rg2vgrxnm.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '2', '2021-11-19 21:37:57', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('193', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftsrl7r9nnngyv4ku.jpg', '/public/uploads/2021-11-19/cftsrl7r9nnngyv4ku.jpg', 'cftsrl7r9nnngyv4ku.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '2', '2021-11-19 21:40:14', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('194', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg', '/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg', 'cftsrmz4obqsgk9ndy.jpg', '8285710645fd9ca7.jpg', '81341', '1', '.jpg', '2', '2021-11-19 21:40:18', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('195', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', '/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', 'cftswh0hteh6uoswhh.png', 'vip.png', '1116', '1', '.png', '2', '2021-11-19 21:46:37', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('196', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3pn460xj68z3oqx2.jpg', '/public/uploads/2021-12-01/cg3pn460xj68z3oqx2.jpg', 'cg3pn460xj68z3oqx2.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 13:20:00', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('197', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3pnen8hj9gjj8xdj.jpg', '/public/uploads/2021-12-01/cg3pnen8hj9gjj8xdj.jpg', 'cg3pnen8hj9gjj8xdj.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 13:20:23', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('198', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3po0lljd1kiz4v4m.jpg', '/public/uploads/2021-12-01/cg3po0lljd1kiz4v4m.jpg', 'cg3po0lljd1kiz4v4m.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 13:21:10', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('199', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3poxw2vglk3xa5eu.jpg', '/public/uploads/2021-12-01/cg3poxw2vglk3xa5eu.jpg', 'cg3poxw2vglk3xa5eu.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 13:22:23', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('200', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3poxwauhe4vqzpfz.jpg', '/public/uploads/2021-12-01/cg3poxwauhe4vqzpfz.jpg', 'cg3poxwauhe4vqzpfz.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '1', '2021-12-01 13:22:23', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('201', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3poxwh4k4kqlckzh.jpg', '/public/uploads/2021-12-01/cg3poxwh4k4kqlckzh.jpg', 'cg3poxwh4k4kqlckzh.jpg', '20211010-JsiZFi.jpg', '18953', '1', '.jpg', '1', '2021-12-01 13:22:23', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('202', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3pphtk70twiodilm.jpg', '/public/uploads/2021-12-01/cg3pphtk70twiodilm.jpg', 'cg3pphtk70twiodilm.jpg', '8285710645fd9ca7.jpg', '81341', '1', '.jpg', '1', '2021-12-01 13:23:06', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('203', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3swz1tf03gsvxkw4.jpg', '/public/uploads/2021-12-01/cg3swz1tf03gsvxkw4.jpg', 'cg3swz1tf03gsvxkw4.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 15:53:56', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('204', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t1h5zm9sspyew5h.jpg', '/public/uploads/2021-12-01/cg3t1h5zm9sspyew5h.jpg', 'cg3t1h5zm9sspyew5h.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 15:59:48', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('205', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t1w35304cmuwucw.jpg', '/public/uploads/2021-12-01/cg3t1w35304cmuwucw.jpg', 'cg3t1w35304cmuwucw.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 16:00:21', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('206', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t1w3cyglgqponc8.jpg', '/public/uploads/2021-12-01/cg3t1w3cyglgqponc8.jpg', 'cg3t1w3cyglgqponc8.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 16:00:21', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('207', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t1w3joi94rqfffn.jpg', '/public/uploads/2021-12-01/cg3t1w3joi94rqfffn.jpg', 'cg3t1w3joi94rqfffn.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '1', '2021-12-01 16:00:21', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('208', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2aeeg0hwbpjpv7.jpg', '/public/uploads/2021-12-01/cg3t2aeeg0hwbpjpv7.jpg', 'cg3t2aeeg0hwbpjpv7.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 16:00:52', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('209', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2aem9yf4hxl2j6.jpg', '/public/uploads/2021-12-01/cg3t2aem9yf4hxl2j6.jpg', 'cg3t2aem9yf4hxl2j6.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 16:00:52', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('210', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2aes7yl4bmxcc1.jpg', '/public/uploads/2021-12-01/cg3t2aes7yl4bmxcc1.jpg', 'cg3t2aes7yl4bmxcc1.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '1', '2021-12-01 16:00:52', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('211', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2aez270k0wjf47.jpg', '/public/uploads/2021-12-01/cg3t2aez270k0wjf47.jpg', 'cg3t2aez270k0wjf47.jpg', '20211010-JsiZFi.jpg', '18953', '1', '.jpg', '1', '2021-12-01 16:00:52', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('212', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2v3eax84r6avqj.jpg', '/public/uploads/2021-12-01/cg3t2v3eax84r6avqj.jpg', 'cg3t2v3eax84r6avqj.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 16:01:37', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('213', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2v3o7yp4ai9xrg.jpg', '/public/uploads/2021-12-01/cg3t2v3o7yp4ai9xrg.jpg', 'cg3t2v3o7yp4ai9xrg.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 16:01:37', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('214', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t2v3w719klrdtlc.jpg', '/public/uploads/2021-12-01/cg3t2v3w719klrdtlc.jpg', 'cg3t2v3w719klrdtlc.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '1', '2021-12-01 16:01:37', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('215', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t3cw2qr44n3eh5u.jpg', '/public/uploads/2021-12-01/cg3t3cw2qr44n3eh5u.jpg', 'cg3t3cw2qr44n3eh5u.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 16:02:16', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('216', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t3cwadgxgwtnxip.jpg', '/public/uploads/2021-12-01/cg3t3cwadgxgwtnxip.jpg', 'cg3t3cwadgxgwtnxip.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 16:02:16', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('217', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t3cwgz4eosnmqpf.jpg', '/public/uploads/2021-12-01/cg3t3cwgz4eosnmqpf.jpg', 'cg3t3cwgz4eosnmqpf.jpg', '20210906-VdEor9.jpg', '48769', '1', '.jpg', '1', '2021-12-01 16:02:16', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('218', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3t3cwnv9cs1sytgu.jpg', '/public/uploads/2021-12-01/cg3t3cwnv9cs1sytgu.jpg', 'cg3t3cwnv9cs1sytgu.jpg', '20211010-JsiZFi.jpg', '18953', '1', '.jpg', '1', '2021-12-01 16:02:16', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('219', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3tlw6akye8q7n1dd.jpg', '/public/uploads/2021-12-01/cg3tlw6akye8q7n1dd.jpg', 'cg3tlw6akye8q7n1dd.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 16:26:28', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('220', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3ukmgziz1wjummpn.jpg', '/public/uploads/2021-12-01/cg3ukmgziz1wjummpn.jpg', 'cg3ukmgziz1wjummpn.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 17:11:50', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('221', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3ul188hrggp5oouq.jpg', '/public/uploads/2021-12-01/cg3ul188hrggp5oouq.jpg', 'cg3ul188hrggp5oouq.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-01 17:12:22', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('222', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3ul3xs13b4fkfcfl.jpg', '/public/uploads/2021-12-01/cg3ul3xs13b4fkfcfl.jpg', 'cg3ul3xs13b4fkfcfl.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-01 17:12:28', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('223', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3umr4plipwufw5be.jpg', '/public/uploads/2021-12-01/cg3umr4plipwufw5be.jpg', 'cg3umr4plipwufw5be.jpg', '20211011-OARliX.jpg', '20999', '1', '.jpg', '1', '2021-12-01 17:14:37', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('224', '1', 'http://localhost:8199/public/uploads/2021-12-01/cg3uqqylnfccaipjrs.jpg', '/public/uploads/2021-12-01/cg3uqqylnfccaipjrs.jpg', 'cg3uqqylnfccaipjrs.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '2', '2021-12-01 17:19:50', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('225', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5fnzxbhaikhynlvt.jpg', '/public/uploads/2021-12-03/cg5fnzxbhaikhynlvt.jpg', 'cg5fnzxbhaikhynlvt.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2021-12-03 13:56:18', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('226', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5fo5aw3gwkuxncmb.jpg', '/public/uploads/2021-12-03/cg5fo5aw3gwkuxncmb.jpg', 'cg5fo5aw3gwkuxncmb.jpg', '20211026-NOnYqA.jpg', '10657', '1', '.jpg', '1', '2021-12-03 13:56:29', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('227', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5fvnup07tws9u7eb.jpg', '/public/uploads/2021-12-03/cg5fvnup07tws9u7eb.jpg', 'cg5fvnup07tws9u7eb.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2021-12-03 14:06:18', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('228', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5fz5zjb4skh1ggup.jpg', '/public/uploads/2021-12-03/cg5fz5zjb4skh1ggup.jpg', 'cg5fz5zjb4skh1ggup.jpg', '8285710645fd9ca7.jpg', '81341', '1', '.jpg', '1', '2021-12-03 14:10:53', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('229', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5fzyh7pjscv5syem.jpg', '/public/uploads/2021-12-03/cg5fzyh7pjscv5syem.jpg', 'cg5fzyh7pjscv5syem.jpg', '8285710645fd9ca7.jpg', '81341', '1', '.jpg', '1', '2021-12-03 14:11:55', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('230', '1', 'http://localhost:8199/public/uploads/2021-12-03/cg5g09rq20o8zbigue.jpg', '/public/uploads/2021-12-03/cg5g09rq20o8zbigue.jpg', 'cg5g09rq20o8zbigue.jpg', '20211026-NOnYqA.jpg', '10657', '1', '.jpg', '1', '2021-12-03 14:12:19', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('231', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfxx0pv3wfkzju3am.png', '/public/uploads/2021-12-15/cgfxx0pv3wfkzju3am.png', 'cgfxx0pv3wfkzju3am.png', 'pm2.png', '54766', '1', '.png', '2', '2021-12-15 22:21:04', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('232', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfxyb1fyqx052bryo.png', '/public/uploads/2021-12-15/cgfxyb1fyqx052bryo.png', 'cgfxyb1fyqx052bryo.png', 'Supervisor.png', '7751', '1', '.png', '2', '2021-12-15 22:22:45', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('233', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfy0qxwfflg4jyjb3.png', '/public/uploads/2021-12-15/cgfy0qxwfflg4jyjb3.png', 'cgfy0qxwfflg4jyjb3.png', 'redis.png', '13299', '1', '.png', '2', '2021-12-15 22:25:56', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('234', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfy6d37qw5gcsmleu.png', '/public/uploads/2021-12-15/cgfy6d37qw5gcsmleu.png', 'cgfy6d37qw5gcsmleu.png', 'createweb.png', '47980', '1', '.png', '2', '2021-12-15 22:33:16', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('235', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfy6iv1k1ugstsy2p.png', '/public/uploads/2021-12-15/cgfy6iv1k1ugstsy2p.png', 'cgfy6iv1k1ugstsy2p.png', 'createweb.png', '47980', '1', '.png', '2', '2021-12-15 22:33:29', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('236', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfyj4hjn9dopnx8xk.png', '/public/uploads/2021-12-15/cgfyj4hjn9dopnx8xk.png', 'cgfyj4hjn9dopnx8xk.png', 'createf.png', '34666', '1', '.png', '1', '2021-12-15 22:49:56', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('237', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfyq4o3u4soshk4u3.png', '/public/uploads/2021-12-15/cgfyq4o3u4soshk4u3.png', 'cgfyq4o3u4soshk4u3.png', 'updataserve.png', '41447', '1', '.png', '1', '2021-12-15 22:59:05', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('238', '1', 'http://localhost:8199/public/uploads/2021-12-15/cgfyqun89l10nm8unb.png', '/public/uploads/2021-12-15/cgfyqun89l10nm8unb.png', 'cgfyqun89l10nm8unb.png', 'serve上传完成.png', '24120', '1', '.png', '1', '2021-12-15 23:00:02', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('239', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0bl36oa9wznf9wt.png', '/public/uploads/2021-12-16/cgg0bl36oa9wznf9wt.png', 'cgg0bl36oa9wznf9wt.png', '上传完成.png', '44340', '1', '.png', '2', '2021-12-16 00:14:08', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('240', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0d65k2za0lkz5te.png', '/public/uploads/2021-12-16/cgg0d65k2za0lkz5te.png', 'cgg0d65k2za0lkz5te.png', '解压.png', '51730', '1', '.png', '2', '2021-12-16 00:16:12', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('241', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0i5dau82kqyidme.png', '/public/uploads/2021-12-16/cgg0i5dau82kqyidme.png', 'cgg0i5dau82kqyidme.png', 'config.png', '118569', '1', '.png', '2', '2021-12-16 00:22:42', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('242', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0i97deaw8f3wsmv.png', '/public/uploads/2021-12-16/cgg0i97deaw8f3wsmv.png', 'cgg0i97deaw8f3wsmv.png', 'config2.png', '31253', '1', '.png', '2', '2021-12-16 00:22:51', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('243', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0nmrnc18kavjwkr.png', '/public/uploads/2021-12-16/cgg0nmrnc18kavjwkr.png', 'cgg0nmrnc18kavjwkr.png', 'redis.png', '178509', '1', '.png', '2', '2021-12-16 00:29:52', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('244', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0ojf7guiwxua8xl.png', '/public/uploads/2021-12-16/cgg0ojf7guiwxua8xl.png', 'cgg0ojf7guiwxua8xl.png', '导入数据.png', '31308', '1', '.png', '1', '2021-12-16 00:31:03', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('245', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0p83zmumgqzfz2a.png', '/public/uploads/2021-12-16/cgg0p83zmumgqzfz2a.png', 'cgg0p83zmumgqzfz2a.png', '数据库.png', '73957', '1', '.png', '1', '2021-12-16 00:31:57', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('246', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0snaj324cgfnbzo.png', '/public/uploads/2021-12-16/cgg0snaj324cgfnbzo.png', 'cgg0snaj324cgfnbzo.png', 'ng.png', '99043', '1', '.png', '2', '2021-12-16 00:36:25', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('247', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg0urto8j54gflchd.png', '/public/uploads/2021-12-16/cgg0urto8j54gflchd.png', 'cgg0urto8j54gflchd.png', 'pm2.png', '161389', '1', '.png', '2', '2021-12-16 00:39:11', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('248', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg13i5f0sgkkm4wqs.png', '/public/uploads/2021-12-16/cgg13i5f0sgkkm4wqs.png', 'cgg13i5f0sgkkm4wqs.png', 'fb.png', '157060', '1', '.png', '2', '2021-12-16 00:50:36', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('249', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg175jvd1ekfsezbk.png', '/public/uploads/2021-12-16/cgg175jvd1ekfsezbk.png', 'cgg175jvd1ekfsezbk.png', 'ip.png', '12944', '1', '.png', '1', '2021-12-16 00:55:22', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('250', '1', 'http://localhost:8199/public/uploads/2021-12-16/cgg17nsvnkmgwz7ux5.png', '/public/uploads/2021-12-16/cgg17nsvnkmgwz7ux5.png', 'cgg17nsvnkmgwz7ux5.png', 'ip.png', '12944', '1', '.png', '2', '2021-12-16 00:56:01', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('275', '1', 'http://localhost:8199/public/uploads/2021-12-18/cgiiy4qqaefcaceapb.png', '/public/uploads/2021-12-18/cgiiy4qqaefcaceapb.png', 'cgiiy4qqaefcaceapb.png', 'logo.png', '4000', '1', '.png', '2', '2021-12-18 23:15:15', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('276', '1', 'http://localhost:8199/public/uploads/2021-12-20/cgk8y59ho7a8u1gnqs.mp3', '/public/uploads/2021-12-20/cgk8y59ho7a8u1gnqs.mp3', 'cgk8y59ho7a8u1gnqs.mp3', '8477751_1623867216.mp3', '3683175', '1', '.mp3', '1', '2021-12-20 23:50:25', null, 'MP3', '');
INSERT INTO `sys_media` VALUES ('277', '1', 'http://localhost:8199/public/uploads/2021-12-30/cgsmtb85p828pdscxz.png', '/public/uploads/2021-12-30/cgsmtb85p828pdscxz.png', 'cgsmtb85p828pdscxz.png', 'logo (2).png', '6573', '1', '.png', '2', '2021-12-30 20:23:41', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('278', '1', 'http://localhost:8199/public/uploads/2021-12-30/cgsmufeakogxomn2wq.png', '/public/uploads/2021-12-30/cgsmufeakogxomn2wq.png', 'cgsmufeakogxomn2wq.png', 'logo (2).png', '6573', '1', '.png', '2', '2021-12-30 20:25:08', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('279', '1', 'http://localhost:8199/public/uploads/2021-12-30/cgsr33kpzlw5euvozb.jpg', '/public/uploads/2021-12-30/cgsr33kpzlw5euvozb.jpg', 'cgsr33kpzlw5euvozb.jpg', '1626944000-mishanwu.jpg', '149512', '1', '.jpg', '1', '2021-12-30 23:44:32', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('280', '1', 'http://krxun.com/public/uploads/2022-01-11/ch2wy8uosrwox35jmy.png', '/public/uploads/2022-01-11/ch2wy8uosrwox35jmy.png', 'ch2wy8uosrwox35jmy.png', 'createWeb.png', '47084', '1', '.png', '1', '2022-01-11 22:26:58', null, 'PNG', '');
INSERT INTO `sys_media` VALUES ('281', '1', 'http://krxun.com/public/uploads/2022-03-10/cig3fx3j1oco0ajpen.jpg', '/public/uploads/2022-03-10/cig3fx3j1oco0ajpen.jpg', 'cig3fx3j1oco0ajpen.jpg', '3e138bbc65502095.jpg', '58778', '1', '.jpg', '1', '2022-03-10 17:52:48', null, 'JPG', '');
INSERT INTO `sys_media` VALUES ('282', '1', 'http://krxun.com/public/uploads/2022-03-12/cihx7jlm4958isqp6d.jpg', '/public/uploads/2022-03-12/cihx7jlm4958isqp6d.jpg', 'cihx7jlm4958isqp6d.jpg', '80bd6c8d9f411e07.jpg', '55549', '1', '.jpg', '1', '2022-03-12 21:25:05', null, 'JPG', '');

-- ----------------------------
-- Table structure for sys_media_related
-- ----------------------------
DROP TABLE IF EXISTS `sys_media_related`;
CREATE TABLE `sys_media_related` (
  `media_id` bigint(20) NOT NULL,
  `related_id` bigint(20) NOT NULL,
  `module` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_media_related
-- ----------------------------
INSERT INTO `sys_media_related` VALUES ('164', '4', 'resource');
INSERT INTO `sys_media_related` VALUES ('167', '3', 'resource');
INSERT INTO `sys_media_related` VALUES ('168', '2', 'resource');
INSERT INTO `sys_media_related` VALUES ('169', '5', 'resource');
INSERT INTO `sys_media_related` VALUES ('170', '6', 'resource');
INSERT INTO `sys_media_related` VALUES ('171', '7', 'resource');
INSERT INTO `sys_media_related` VALUES ('172', '2', 'video');
INSERT INTO `sys_media_related` VALUES ('173', '2', 'video');
INSERT INTO `sys_media_related` VALUES ('173', '3', 'video');
INSERT INTO `sys_media_related` VALUES ('174', '3', 'video');
INSERT INTO `sys_media_related` VALUES ('173', '4', 'video');
INSERT INTO `sys_media_related` VALUES ('175', '4', 'video');
INSERT INTO `sys_media_related` VALUES ('173', '5', 'video');
INSERT INTO `sys_media_related` VALUES ('176', '5', 'video');
INSERT INTO `sys_media_related` VALUES ('177', '3', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '3', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '2', 'audio');
INSERT INTO `sys_media_related` VALUES ('179', '2', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '4', 'audio');
INSERT INTO `sys_media_related` VALUES ('181', '4', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '5', 'audio');
INSERT INTO `sys_media_related` VALUES ('182', '5', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '6', 'audio');
INSERT INTO `sys_media_related` VALUES ('183', '6', 'audio');
INSERT INTO `sys_media_related` VALUES ('178', '7', 'audio');
INSERT INTO `sys_media_related` VALUES ('184', '7', 'audio');
INSERT INTO `sys_media_related` VALUES ('185', '2', 'edu');
INSERT INTO `sys_media_related` VALUES ('186', '3', 'edu');
INSERT INTO `sys_media_related` VALUES ('187', '4', 'edu');
INSERT INTO `sys_media_related` VALUES ('188', '5', 'edu');
INSERT INTO `sys_media_related` VALUES ('189', '2', 'article');
INSERT INTO `sys_media_related` VALUES ('190', '3', 'article');
INSERT INTO `sys_media_related` VALUES ('191', '1', 'article');
INSERT INTO `sys_media_related` VALUES ('192', '4', 'article');
INSERT INTO `sys_media_related` VALUES ('195', '1', 'vip');
INSERT INTO `sys_media_related` VALUES ('195', '2', 'vip');
INSERT INTO `sys_media_related` VALUES ('195', '3', 'vip');
INSERT INTO `sys_media_related` VALUES ('195', '4', 'vip');
INSERT INTO `sys_media_related` VALUES ('224', '16', 'topic');
INSERT INTO `sys_media_related` VALUES ('193', '1', 'user');
INSERT INTO `sys_media_related` VALUES ('194', '1', 'user');
INSERT INTO `sys_media_related` VALUES ('231', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('232', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('233', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('234', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('235', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('239', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('240', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('241', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('242', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('243', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('246', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('247', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('248', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('250', '5', 'article');
INSERT INTO `sys_media_related` VALUES ('278', '0', 'BaseSetting');

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `notice_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `from_user_id` bigint(20) DEFAULT NULL,
  `receiver` bigint(20) DEFAULT NULL COMMENT '接收者',
  `detail_module` varchar(50) DEFAULT NULL,
  `detail_id` bigint(20) DEFAULT NULL,
  `system_type` tinyint(4) DEFAULT NULL COMMENT '1用户注册，2内容删除，3打赏收入，4购买收入，报名收入',
  `type` tinyint(4) NOT NULL COMMENT '通知类型 1用户注册通知,2内容删除通知,3用户收入通知,4评论通知,5回答通知,6获赞通知,7收到关注通知',
  `content` text COMMENT '内容',
  `status` tinyint(4) DEFAULT '0' COMMENT '是否阅读1未阅读，2已阅读',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
INSERT INTO `sys_notice` VALUES ('101', '1', '1', 'comment', '0', '0', '2', 's撒大噶高大上的噶啥的噶', '1', '2022-03-12 17:57:04');
INSERT INTO `sys_notice` VALUES ('102', '1', '1', 'comment', '0', '0', '2', '大发噶收到法国撒旦发个', '1', '2022-03-12 17:57:06');
INSERT INTO `sys_notice` VALUES ('103', '1', '1', 'comment', '107', '0', '2', '手动阀手动阀撒旦', '2', '2022-03-12 17:57:08');
INSERT INTO `sys_notice` VALUES ('104', '1', '1', 'comment', '108', '0', '2', '十大发射点发射点发射点', '2', '2022-03-12 17:57:12');
INSERT INTO `sys_notice` VALUES ('105', '1', '1', 'comment', '106', '0', '2', '的风格和对方给', '2', '2022-03-12 19:16:37');
INSERT INTO `sys_notice` VALUES ('106', '1', '1', 'comment', '109', '0', '4', '点赞了你发布的《十大发射点发射点发射点》评论', '2', '2022-03-12 19:18:17');
INSERT INTO `sys_notice` VALUES ('107', '1', '1', 'topic', '20', '0', '4', '点赞了你发布的《岁的法国撒的分公司的还是的风格和》帖子', '2', '2022-03-12 20:59:37');
INSERT INTO `sys_notice` VALUES ('108', '1', '1', 'topic', '25', '0', '4', '点赞了你发布的《的说法伽师的噶SDF士大夫》帖子', '2', '2022-03-14 17:17:34');
INSERT INTO `sys_notice` VALUES ('109', '0', '1', 'answer', '28', '0', '3', '撒旦发射点发撒地方', '1', '2022-03-15 12:49:24');
INSERT INTO `sys_notice` VALUES ('110', '1', '1', 'comment', '0', '0', '2', 'fgdhdsfghdfgh', '1', '2022-03-15 14:44:53');
INSERT INTO `sys_notice` VALUES ('111', '1', '1', 'comment', '111', '0', '4', '点赞了你发布的《fgdhdsfghdfgh》评论', '2', '2022-03-15 20:41:43');
INSERT INTO `sys_notice` VALUES ('112', '1', '1', 'article', '5', '0', '4', '点赞了你发布的《Fiber宝塔安装教程》文章', '2', '2022-03-16 18:42:54');
INSERT INTO `sys_notice` VALUES ('113', '0', '1', 'answer', '29', '0', '3', '的事发生的故事知道', '1', '2022-03-31 00:33:35');
INSERT INTO `sys_notice` VALUES ('114', '0', '1', 'answer', '30', '0', '3', '大发噶收到法国撒旦发个', '1', '2022-03-31 00:33:38');
INSERT INTO `sys_notice` VALUES ('115', '0', '1', 'answer', '31', '0', '3', '的方式股市大幅改善大哥', '1', '2022-03-31 00:33:42');
INSERT INTO `sys_notice` VALUES ('116', '0', '1', 'answer', '32', '0', '3', '豆腐干岁的法国撒地方', '1', '2022-03-31 00:33:44');
INSERT INTO `sys_notice` VALUES ('117', '0', '1', 'answer', '33', '0', '3', 'fg士大夫敢死队风格还是的风格和对方给好的', '1', '2022-03-31 00:33:47');
INSERT INTO `sys_notice` VALUES ('118', '0', '1', 'answer', '34', '0', '3', '豆腐干岁的法国大根深蒂固岁的法国士大夫', '1', '2022-03-31 00:33:56');
INSERT INTO `sys_notice` VALUES ('119', '0', '1', 'answer', '35', '0', '3', '地方很多法国还是发给', '1', '2022-03-31 00:35:02');
INSERT INTO `sys_notice` VALUES ('120', '0', '1', 'answer', '36', '0', '3', '大哥的身份噶士大夫公司', '1', '2022-03-31 00:36:06');

-- ----------------------------
-- Table structure for sys_order
-- ----------------------------
DROP TABLE IF EXISTS `sys_order`;
CREATE TABLE `sys_order` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `user_id` bigint(20) DEFAULT NULL COMMENT '下单人id',
  `author_id` bigint(20) DEFAULT NULL COMMENT '作者id',
  `pay_method` tinyint(4) DEFAULT NULL COMMENT '支付方式：1支付宝，2微信，3余额',
  `order_money` decimal(10,2) DEFAULT NULL COMMENT '订单金额',
  `district_money` decimal(10,2) DEFAULT NULL COMMENT '优惠金额',
  `author_money` decimal(10,2) DEFAULT NULL COMMENT '作者收益',
  `service_money` decimal(10,2) DEFAULT NULL COMMENT '服务费',
  `payment_money` decimal(10,2) DEFAULT NULL COMMENT '支付金额',
  `order_point` bigint(10) unsigned DEFAULT '0' COMMENT '订单积分',
  `order_type` tinyint(4) DEFAULT NULL COMMENT '订单类型: 1 充值，2打赏充电，3内容购买，4查看话题隐藏内容  ,5加入付费圈子,6购买付费课程,7查看付费答案，8开通vip，9认证付费',
  `order_mode` tinyint(4) DEFAULT NULL COMMENT '订单方式：1虚拟物品，2实体物品',
  `shipping_money` decimal(10,2) DEFAULT NULL COMMENT '运输费用',
  `shipping_address` varchar(100) DEFAULT NULL COMMENT '收获地址',
  `shipping_comp_name` varchar(10) DEFAULT NULL COMMENT '快递公司',
  `shipping_phone` varchar(50) DEFAULT NULL COMMENT '收货人联系方式',
  `shipping_name` varchar(50) DEFAULT NULL COMMENT '收货人姓名',
  `shipping_sn` varchar(60) DEFAULT NULL COMMENT '快递单号',
  `shipping_time` datetime DEFAULT NULL COMMENT '发货时间',
  `detail_id` bigint(20) DEFAULT NULL,
  `detail_module` varchar(50) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0' COMMENT '订单状态 1未支付，2已支付',
  `invoice` varchar(100) DEFAULT NULL COMMENT '发票抬头',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `receive_time` datetime DEFAULT NULL COMMENT '收货时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '最后修改时间',
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_order
-- ----------------------------
INSERT INTO `sys_order` VALUES ('2', 'C2163690200528265618233', '3', '1', '3', '20.00', '0.00', '19.00', '1.00', '20.00', '0', '2', '1', '0.00', '', '', '', '', '', null, '1', 'user', '2', '', '2021-11-14 23:00:05', null, '2021-11-14 23:00:05', '2021-11-14 23:00:05');
INSERT INTO `sys_order` VALUES ('5', 'Z3163691082650968032333', '3', '1', '3', '20.00', '10.00', '9.50', '0.50', '10.00', '0', '3', '1', '0.00', '', '', '', '', '', null, '2', 'audio', '2', '', '2021-11-15 01:27:06', null, '2021-11-15 01:27:06', '2021-11-15 01:27:06');

-- ----------------------------
-- Table structure for sys_question
-- ----------------------------
DROP TABLE IF EXISTS `sys_question`;
CREATE TABLE `sys_question` (
  `question_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `group_id` bigint(20) DEFAULT NULL,
  `title` varchar(120) DEFAULT NULL,
  `content` text,
  `hots` bigint(20) DEFAULT NULL,
  `favorites` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `anonymous` tinyint(4) DEFAULT NULL COMMENT '是否匿名',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`question_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_question
-- ----------------------------
INSERT INTO `sys_question` VALUES ('1', '1', '2', '守望先锋什么时候凉', '<p>韩国选手辱华，是不是守望先锋就要凉凉</p>', '0', '0', '0', '26', '2', '2', '2021-05-11 06:08:14', '2021-05-11 06:08:14', null, '啊实打实');
INSERT INTO `sys_question` VALUES ('2', '3', '2', '测试问题内按实际亏损扩大覆盖就开始的风格阿松大', '<p>岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大岁的法国撒旦发个士大夫阿松大</p>', '0', '0', '0', '43', '2', '2', '2021-07-13 23:48:57', '2021-07-13 23:48:57', null, '');
INSERT INTO `sys_question` VALUES ('3', '1', '2', '测试问题的是否缴纳', '<p>豆腐干士大夫嘎斯的嘎斯豆腐干阿斯顿噶啥的噶啥的是D</p>', '0', '0', '0', '21', '2', '2', '2021-07-29 01:23:36', '2021-07-29 01:23:36', null, '');
INSERT INTO `sys_question` VALUES ('4', '1', '2', '测试新问题对法轮功空间的是否客户及时反馈过来', '<p>高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬高峰会上发给的活动分工和是否更换士大夫gadfly阿斯蒂芬阿斯蒂芬</p>', '0', '0', '0', '18', '1', '2', '2021-08-24 01:11:42', '2021-08-24 01:11:42', null, '');
INSERT INTO `sys_question` VALUES ('5', '1', '2', '测去啊是的士大夫', '<p><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em><em>地方撒旦发射点风格</em></p>', '0', '0', '1', '23', '2', '2', '2021-10-28 15:46:54', '2021-10-28 15:46:54', null, '');

-- ----------------------------
-- Table structure for sys_recharge
-- ----------------------------
DROP TABLE IF EXISTS `sys_recharge`;
CREATE TABLE `sys_recharge` (
  `recharge_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `code` varchar(50) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL COMMENT '充值金额',
  `mode` tinyint(4) DEFAULT NULL COMMENT '充值方式：1支付宝，2微信，3卡密，4人工转账',
  `card_key` varchar(255) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL COMMENT '转账人名称',
  `type` tinyint(4) DEFAULT NULL COMMENT '转账类型：1支付宝，2微信',
  `number` varchar(50) DEFAULT NULL COMMENT '转账单号',
  `status` tinyint(4) DEFAULT NULL COMMENT '充值状态：1待审核，2已充值，3未充值',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`recharge_id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_recharge
-- ----------------------------
INSERT INTO `sys_recharge` VALUES ('4', '1', 'C163575685130179567111', '20.00', '1', '', '', '0', '', '1', '2021-11-01 16:54:11', '2021-11-01 16:54:11', null);
INSERT INTO `sys_recharge` VALUES ('5', '1', 'C163575690830426603111', '20.00', '1', '', '', '0', '', '1', '2021-11-01 16:55:08', '2021-11-01 16:55:08', null);
INSERT INTO `sys_recharge` VALUES ('6', '1', 'C163575693700294115511', '20.00', '1', '', '', '0', '', '1', '2021-11-01 16:55:37', '2021-11-01 16:55:37', null);
INSERT INTO `sys_recharge` VALUES ('7', '1', 'C163575696337909285311', '20.00', '1', '', '', '0', '', '1', '2021-11-01 16:56:03', '2021-11-01 16:56:03', null);
INSERT INTO `sys_recharge` VALUES ('8', '1', 'C163575710798931870311', '20.00', '1', '', '', '0', '', '1', '2021-11-01 16:58:27', '2021-11-01 16:58:27', null);
INSERT INTO `sys_recharge` VALUES ('9', '1', 'C163575988003100317511', '20.00', '1', '', '', '0', '', '1', '2021-11-01 17:44:40', '2021-11-01 17:44:40', null);
INSERT INTO `sys_recharge` VALUES ('10', '1', 'C163575989939443448411', '20.00', '1', '', '', '0', '', '1', '2021-11-01 17:44:59', '2021-11-01 17:44:59', null);
INSERT INTO `sys_recharge` VALUES ('11', '1', 'C163575995593712354911', '20.00', '1', '', '', '0', '', '1', '2021-11-01 17:45:55', '2021-11-01 17:45:55', null);
INSERT INTO `sys_recharge` VALUES ('12', '1', 'C163575999898850781711', '20.00', '4', '', '123123', '1', '12312', '1', '2021-11-01 17:46:38', '2021-11-01 17:46:38', null);
INSERT INTO `sys_recharge` VALUES ('27', '1', 'C163577563391278250741', '500.00', '4', 'rdfghdsfhsfdgh', 'fgdfg', '1', '45322341123', '2', '2021-11-01 22:07:13', '2021-11-01 22:07:13', '充值成功');
INSERT INTO `sys_recharge` VALUES ('30', '1', 'C163577607835541001141', '0.00', '4', '', 'asdasd', '1', 'asdasd', '2', '2021-11-01 22:14:38', '2021-11-01 22:14:38', '充值成功');

-- ----------------------------
-- Table structure for sys_report
-- ----------------------------
DROP TABLE IF EXISTS `sys_report`;
CREATE TABLE `sys_report` (
  `report_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `related_id` bigint(20) DEFAULT NULL,
  `module` varchar(50) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL COMMENT '举报类型 1广告垃圾，2违规内容，3恶意灌水，4重复发帖',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `remark` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '状态，2 已处理，1 未处理',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`report_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_report
-- ----------------------------
INSERT INTO `sys_report` VALUES ('1', '1', '1', 'topic', '1', 'sdfs', 'sdf', '1', '2021-09-26 17:20:54', null);

-- ----------------------------
-- Table structure for sys_resource
-- ----------------------------
DROP TABLE IF EXISTS `sys_resource`;
CREATE TABLE `sys_resource` (
  `resource_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '发布的用户',
  `cate_id` bigint(20) DEFAULT NULL COMMENT '分类id',
  `hots` bigint(20) DEFAULT NULL,
  `favorites` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL COMMENT '阅读量',
  `title` varchar(150) DEFAULT NULL COMMENT '标题',
  `content` mediumtext COMMENT '内容',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `has_down` tinyint(4) DEFAULT '1' COMMENT '是否有下载1没有，2有',
  `down_mode` tinyint(4) DEFAULT NULL COMMENT '下载权限 0公开下载，1付费下载，2评论下载，3登录下载',
  `price` decimal(10,2) DEFAULT NULL,
  `down_url` text COMMENT '网盘地址',
  `purpose` text COMMENT '用途',
  `attribute` text COMMENT '属性',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_resource
-- ----------------------------
INSERT INTO `sys_resource` VALUES ('2', '1', '5', '21312', '12312', '312312', '336', 'Procreate', '<p>的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方的首发式地方</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftn8mjwg6t9iaomlf.png', '2', '3', '0.00', '[{\"key\":\"fibercms.com\",\"title\":\"百度网盘\",\"val\":\"sdf\"}]', '[{\"key\":\"士大夫\",\"val\":\"sd发\"}]', '[{\"key\":\"sdf\",\"val\":\"士大夫\"}]', '国会尽快汇款给火箭', '2', '2021-07-12 16:59:11', '2021-11-19 17:22:50', null, '');
INSERT INTO `sys_resource` VALUES ('3', '1', '5', '21312', '312312', '3123123', '178', 'snipaste', '<p>测试资源法国空军发给客户发大哥</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftn886w6pso93otxw.png', '2', '1', '213.00', '[{\"key\":\"https://www.baidu.com/\",\"title\":\"百度网盘\",\"val\":\"asdf\"}]', '[]', '[]', '和对方国家和地方各级地方功能和岁的法国的', '2', '2021-08-25 00:58:54', '2021-11-19 17:22:43', null, '');
INSERT INTO `sys_resource` VALUES ('4', '1', '5', '2312', '312312', '3123', '135', 'PS全家桶', '<p>规划局房管局感觉好多好多风格和对方给梵蒂冈回复</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftlm1x82z1ogbe0bm.png', '1', '0', '0.00', '[]', '[]', '[]', 'f的噶丹是法国', '2', '2021-09-23 16:02:57', '2021-11-19 17:21:37', null, '');
INSERT INTO `sys_resource` VALUES ('5', '1', '5', '21312', '312312', '3123', '131', 'Krita中文版', '<p>Krita 是一款自由开源的免费绘画软件，在 GPL 许可证下发布。它针对手绘用途进行设计，内建多种可定制的笔刷系统，适用于绘制概念美术设计、材质、电影接景、插画和漫画等。它能够绘制位图图像、矢量图形和制作动画，具备完整的色彩管理功能。支持 Windows、Linux 、macOS和Android 平台。软件和全套文档均已自带全中文翻译。 官方标语 自由开源免费 绝无功能限制 绝无商用限制 软件介绍 Krita软件界面 Krita 的软件界面直观易用，不会妨碍你的创作流程。面板可以随意拖放、分组、弹出，调整后的界面可以保存为工作区，可以一键在不同工作区之间切换。键盘快捷键、画布视图修饰键均可定制，还准备了一键切换 PS / SAI 的兼容方案。你可以通过上述调整，让软件完全适应你的使用习惯。 可定制的界面 可显示超过 30 种功能面板 可切换多档亮色或者暗色主题 精心打造的自带笔刷库 Krita笔刷库 Krita 自带了超过 100 种专业的笔刷预设，它们照顾到了各种风格的需要。你可以加载其他笔刷包，也可以观察、调整笔刷的参数，自行学习制作笔刷预设。 手抖修正和笔画平滑 Krita 支持多级笔画防抖和平滑。使用手绘笔刷工具时可在工具选项面板的下拉选单打开。默认的基本平滑可以改善快速绘制线条时的棱角，更高级的防抖可以通过延迟笔画绘制并计算平均值来让线条和压感得到稳定。另有力学笔刷工具，在手绘笔刷的基础上添加了虚拟重量和惯性来画出平滑的笔画。 矢量图形和文字排版 Krita矢量图形和文字 Krita 支持矢量图形，还内建了矢量图形库可以随意拖放素材，方便漫画制作。你可以修改矢量形状的锚点，打造自己的矢量素材库。你可以通过文字工具插入文本并进行简单的排版。Krita 的矢量图形完全兼容 SVG 规范，可以导入或者导出 SVG 图形。 笔刷引擎可深度定制 Krita 内建了 9 种笔刷引擎，每种引擎的建模都针对某种特定功能，如颜色涂抹、轮廓填色、粒子甚至滤镜引擎。每种笔刷引擎带有丰富的定制选项，可以用来打造无数种笔刷预设。笔刷预设可以通过标签来管理和加载。 四方连续和绘画辅助尺 Krita 支持四方连续显示，可在视图菜单中切换。此功能可以让图像在上下左右四个方向连续显示画布内容，实时同步更新，对于制作无缝材质和图案非常有用。绘画辅助尺可以吸附手绘笔刷，方便你画出特定形状或者透视。 标签和资源包管理 Krita标签和资源包管理</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftnb42nzy21h5aw49.png', '1', '0', '0.00', '[]', '[]', '[]', 'Krita 是一款自由开源的免费绘画软件，在 GPL 许可证下发布。它针对手绘用途进行设计，内建多种可定制的笔刷系统，适用于绘制概念美术设计、材质、电影接景、插画和漫画等。它能够绘制位图图像、矢量图形和制作动画，具备完整的色彩管理功能。支持 Windows、Linux 、macO', '2', '2021-11-19 17:24:21', '2021-11-19 17:24:21', null, '');
INSERT INTO `sys_resource` VALUES ('6', '1', '5', '213', '12312', '3123', '1237', 'Artstudio', '<p>的噶士大夫士大夫士大夫</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftnc1l4tz1zml3d8h.png', '1', '0', '0.00', '[]', '[]', '[]', '分公司的发挥风格豆腐干规划局', '2', '2021-11-19 17:25:04', '2021-11-19 17:25:04', null, '');
INSERT INTO `sys_resource` VALUES ('7', '1', '5', '1231231', '3123', '123123', '243', 'LIVE2D学习版', '<p>甚至在媒体开发中也统一工作形象 如今，从原始媒​​体到各种媒体的各种发展并不少见。但是，每种媒体的视觉效果都不同。使用Live2D，您可以使用一种模型开发各种媒体，例如视频，应用程序，游戏等，从而可以统一视觉图像。 使用SDK支持广泛的用途 自2008年推出以来，Live2D已被广泛用于多种用途，与iOS和Android等移动操作系统，家庭视频游戏机以及游戏开发环境Unity兼容。此外，它被广泛用于各种应用程序中，例如将其合并到用于事件和实时视频分发的通信工具中。 您实际上是如何做到的？ 原始图片的准备和处理 原始图像数据是处理的基础 拆卸要移动的每个零件</p>', 'http://localhost:8199/public/uploads/2021-11-19/cftncplplm0law4plg.png', '1', '0', '0.00', '[]', '[]', '[]', '一种被称为LIVE2D的图像表达技术，可以栩栩如生地实现2D图像的立体表达', '2', '2021-11-19 17:26:10', '2021-11-19 17:26:10', null, '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `title` varchar(30) NOT NULL COMMENT '角色名称',
  `status` tinyint(4) NOT NULL COMMENT '角色状态（2正常 1停用）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '管理员', '2', '2021-08-26 11:38:26', '2021-12-17 16:40:49', '');
INSERT INTO `sys_role` VALUES ('15', '后台演示', '2', '2021-08-26 11:41:08', '2021-11-23 20:01:52', '');
INSERT INTO `sys_role` VALUES ('16', '测试角色xxx', '2', '2021-09-20 01:12:54', '2021-09-20 01:13:23', '');

-- ----------------------------
-- Table structure for sys_role_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_authority`;
CREATE TABLE `sys_role_authority` (
  `role_id` bigint(20) NOT NULL,
  `authority_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_authority
-- ----------------------------
INSERT INTO `sys_role_authority` VALUES ('1', '1');
INSERT INTO `sys_role_authority` VALUES ('1', '2');
INSERT INTO `sys_role_authority` VALUES ('1', '3');
INSERT INTO `sys_role_authority` VALUES ('1', '4');
INSERT INTO `sys_role_authority` VALUES ('1', '5');
INSERT INTO `sys_role_authority` VALUES ('1', '8');
INSERT INTO `sys_role_authority` VALUES ('1', '9');
INSERT INTO `sys_role_authority` VALUES ('1', '10');
INSERT INTO `sys_role_authority` VALUES ('1', '11');
INSERT INTO `sys_role_authority` VALUES ('1', '12');
INSERT INTO `sys_role_authority` VALUES ('1', '14');
INSERT INTO `sys_role_authority` VALUES ('1', '15');
INSERT INTO `sys_role_authority` VALUES ('1', '16');
INSERT INTO `sys_role_authority` VALUES ('1', '17');
INSERT INTO `sys_role_authority` VALUES ('1', '18');
INSERT INTO `sys_role_authority` VALUES ('1', '19');
INSERT INTO `sys_role_authority` VALUES ('1', '20');
INSERT INTO `sys_role_authority` VALUES ('1', '21');
INSERT INTO `sys_role_authority` VALUES ('1', '27');
INSERT INTO `sys_role_authority` VALUES ('1', '28');
INSERT INTO `sys_role_authority` VALUES ('1', '29');
INSERT INTO `sys_role_authority` VALUES ('1', '30');
INSERT INTO `sys_role_authority` VALUES ('1', '31');
INSERT INTO `sys_role_authority` VALUES ('1', '32');
INSERT INTO `sys_role_authority` VALUES ('1', '34');
INSERT INTO `sys_role_authority` VALUES ('1', '35');
INSERT INTO `sys_role_authority` VALUES ('1', '36');
INSERT INTO `sys_role_authority` VALUES ('1', '37');
INSERT INTO `sys_role_authority` VALUES ('1', '38');
INSERT INTO `sys_role_authority` VALUES ('1', '39');
INSERT INTO `sys_role_authority` VALUES ('1', '40');
INSERT INTO `sys_role_authority` VALUES ('1', '42');
INSERT INTO `sys_role_authority` VALUES ('1', '43');
INSERT INTO `sys_role_authority` VALUES ('1', '44');
INSERT INTO `sys_role_authority` VALUES ('1', '45');
INSERT INTO `sys_role_authority` VALUES ('1', '46');
INSERT INTO `sys_role_authority` VALUES ('1', '47');
INSERT INTO `sys_role_authority` VALUES ('1', '48');
INSERT INTO `sys_role_authority` VALUES ('1', '52');
INSERT INTO `sys_role_authority` VALUES ('1', '53');
INSERT INTO `sys_role_authority` VALUES ('1', '54');
INSERT INTO `sys_role_authority` VALUES ('1', '55');
INSERT INTO `sys_role_authority` VALUES ('1', '56');
INSERT INTO `sys_role_authority` VALUES ('1', '57');
INSERT INTO `sys_role_authority` VALUES ('1', '58');
INSERT INTO `sys_role_authority` VALUES ('1', '59');
INSERT INTO `sys_role_authority` VALUES ('1', '60');
INSERT INTO `sys_role_authority` VALUES ('1', '61');
INSERT INTO `sys_role_authority` VALUES ('1', '62');
INSERT INTO `sys_role_authority` VALUES ('1', '63');
INSERT INTO `sys_role_authority` VALUES ('1', '64');
INSERT INTO `sys_role_authority` VALUES ('1', '65');
INSERT INTO `sys_role_authority` VALUES ('1', '66');
INSERT INTO `sys_role_authority` VALUES ('1', '67');
INSERT INTO `sys_role_authority` VALUES ('1', '68');
INSERT INTO `sys_role_authority` VALUES ('1', '69');
INSERT INTO `sys_role_authority` VALUES ('1', '70');
INSERT INTO `sys_role_authority` VALUES ('1', '71');
INSERT INTO `sys_role_authority` VALUES ('1', '72');
INSERT INTO `sys_role_authority` VALUES ('1', '73');
INSERT INTO `sys_role_authority` VALUES ('1', '74');
INSERT INTO `sys_role_authority` VALUES ('1', '75');
INSERT INTO `sys_role_authority` VALUES ('1', '76');
INSERT INTO `sys_role_authority` VALUES ('1', '77');
INSERT INTO `sys_role_authority` VALUES ('1', '78');
INSERT INTO `sys_role_authority` VALUES ('1', '79');
INSERT INTO `sys_role_authority` VALUES ('1', '80');
INSERT INTO `sys_role_authority` VALUES ('1', '81');
INSERT INTO `sys_role_authority` VALUES ('1', '82');
INSERT INTO `sys_role_authority` VALUES ('1', '83');
INSERT INTO `sys_role_authority` VALUES ('1', '84');
INSERT INTO `sys_role_authority` VALUES ('1', '85');
INSERT INTO `sys_role_authority` VALUES ('1', '86');
INSERT INTO `sys_role_authority` VALUES ('1', '87');
INSERT INTO `sys_role_authority` VALUES ('1', '88');
INSERT INTO `sys_role_authority` VALUES ('1', '89');
INSERT INTO `sys_role_authority` VALUES ('1', '90');
INSERT INTO `sys_role_authority` VALUES ('1', '91');
INSERT INTO `sys_role_authority` VALUES ('1', '92');
INSERT INTO `sys_role_authority` VALUES ('1', '93');
INSERT INTO `sys_role_authority` VALUES ('1', '94');
INSERT INTO `sys_role_authority` VALUES ('1', '95');
INSERT INTO `sys_role_authority` VALUES ('1', '96');
INSERT INTO `sys_role_authority` VALUES ('1', '97');
INSERT INTO `sys_role_authority` VALUES ('1', '98');
INSERT INTO `sys_role_authority` VALUES ('1', '99');
INSERT INTO `sys_role_authority` VALUES ('1', '100');
INSERT INTO `sys_role_authority` VALUES ('1', '101');
INSERT INTO `sys_role_authority` VALUES ('1', '102');
INSERT INTO `sys_role_authority` VALUES ('1', '103');
INSERT INTO `sys_role_authority` VALUES ('1', '104');
INSERT INTO `sys_role_authority` VALUES ('1', '105');
INSERT INTO `sys_role_authority` VALUES ('1', '106');
INSERT INTO `sys_role_authority` VALUES ('1', '107');
INSERT INTO `sys_role_authority` VALUES ('1', '108');
INSERT INTO `sys_role_authority` VALUES ('1', '109');
INSERT INTO `sys_role_authority` VALUES ('1', '110');
INSERT INTO `sys_role_authority` VALUES ('1', '111');
INSERT INTO `sys_role_authority` VALUES ('1', '112');
INSERT INTO `sys_role_authority` VALUES ('1', '113');
INSERT INTO `sys_role_authority` VALUES ('1', '114');
INSERT INTO `sys_role_authority` VALUES ('1', '115');
INSERT INTO `sys_role_authority` VALUES ('1', '116');
INSERT INTO `sys_role_authority` VALUES ('1', '117');
INSERT INTO `sys_role_authority` VALUES ('1', '118');
INSERT INTO `sys_role_authority` VALUES ('1', '119');
INSERT INTO `sys_role_authority` VALUES ('1', '120');
INSERT INTO `sys_role_authority` VALUES ('1', '121');
INSERT INTO `sys_role_authority` VALUES ('1', '122');
INSERT INTO `sys_role_authority` VALUES ('1', '123');
INSERT INTO `sys_role_authority` VALUES ('1', '124');
INSERT INTO `sys_role_authority` VALUES ('1', '125');
INSERT INTO `sys_role_authority` VALUES ('1', '126');
INSERT INTO `sys_role_authority` VALUES ('1', '127');
INSERT INTO `sys_role_authority` VALUES ('1', '128');
INSERT INTO `sys_role_authority` VALUES ('1', '129');
INSERT INTO `sys_role_authority` VALUES ('1', '130');
INSERT INTO `sys_role_authority` VALUES ('1', '131');
INSERT INTO `sys_role_authority` VALUES ('1', '132');
INSERT INTO `sys_role_authority` VALUES ('1', '133');
INSERT INTO `sys_role_authority` VALUES ('1', '134');
INSERT INTO `sys_role_authority` VALUES ('1', '135');
INSERT INTO `sys_role_authority` VALUES ('1', '136');
INSERT INTO `sys_role_authority` VALUES ('1', '137');
INSERT INTO `sys_role_authority` VALUES ('1', '138');
INSERT INTO `sys_role_authority` VALUES ('1', '139');
INSERT INTO `sys_role_authority` VALUES ('1', '140');
INSERT INTO `sys_role_authority` VALUES ('1', '141');
INSERT INTO `sys_role_authority` VALUES ('1', '142');
INSERT INTO `sys_role_authority` VALUES ('1', '143');
INSERT INTO `sys_role_authority` VALUES ('1', '144');
INSERT INTO `sys_role_authority` VALUES ('1', '145');
INSERT INTO `sys_role_authority` VALUES ('1', '146');
INSERT INTO `sys_role_authority` VALUES ('1', '147');
INSERT INTO `sys_role_authority` VALUES ('1', '148');
INSERT INTO `sys_role_authority` VALUES ('1', '149');
INSERT INTO `sys_role_authority` VALUES ('1', '151');
INSERT INTO `sys_role_authority` VALUES ('1', '152');
INSERT INTO `sys_role_authority` VALUES ('1', '153');
INSERT INTO `sys_role_authority` VALUES ('1', '154');
INSERT INTO `sys_role_authority` VALUES ('1', '155');
INSERT INTO `sys_role_authority` VALUES ('1', '156');
INSERT INTO `sys_role_authority` VALUES ('1', '157');
INSERT INTO `sys_role_authority` VALUES ('1', '158');
INSERT INTO `sys_role_authority` VALUES ('1', '159');
INSERT INTO `sys_role_authority` VALUES ('1', '160');
INSERT INTO `sys_role_authority` VALUES ('1', '161');
INSERT INTO `sys_role_authority` VALUES ('1', '162');
INSERT INTO `sys_role_authority` VALUES ('1', '163');
INSERT INTO `sys_role_authority` VALUES ('15', '1');
INSERT INTO `sys_role_authority` VALUES ('15', '2');
INSERT INTO `sys_role_authority` VALUES ('15', '3');
INSERT INTO `sys_role_authority` VALUES ('15', '4');
INSERT INTO `sys_role_authority` VALUES ('15', '5');
INSERT INTO `sys_role_authority` VALUES ('15', '8');
INSERT INTO `sys_role_authority` VALUES ('15', '9');
INSERT INTO `sys_role_authority` VALUES ('15', '10');
INSERT INTO `sys_role_authority` VALUES ('15', '11');
INSERT INTO `sys_role_authority` VALUES ('15', '12');
INSERT INTO `sys_role_authority` VALUES ('15', '14');
INSERT INTO `sys_role_authority` VALUES ('15', '15');
INSERT INTO `sys_role_authority` VALUES ('15', '16');
INSERT INTO `sys_role_authority` VALUES ('15', '17');
INSERT INTO `sys_role_authority` VALUES ('15', '18');
INSERT INTO `sys_role_authority` VALUES ('15', '19');
INSERT INTO `sys_role_authority` VALUES ('15', '20');
INSERT INTO `sys_role_authority` VALUES ('15', '21');
INSERT INTO `sys_role_authority` VALUES ('15', '27');
INSERT INTO `sys_role_authority` VALUES ('15', '28');
INSERT INTO `sys_role_authority` VALUES ('15', '29');
INSERT INTO `sys_role_authority` VALUES ('15', '30');
INSERT INTO `sys_role_authority` VALUES ('15', '31');
INSERT INTO `sys_role_authority` VALUES ('15', '32');
INSERT INTO `sys_role_authority` VALUES ('15', '34');
INSERT INTO `sys_role_authority` VALUES ('15', '35');
INSERT INTO `sys_role_authority` VALUES ('15', '36');
INSERT INTO `sys_role_authority` VALUES ('15', '37');
INSERT INTO `sys_role_authority` VALUES ('15', '38');
INSERT INTO `sys_role_authority` VALUES ('15', '39');
INSERT INTO `sys_role_authority` VALUES ('15', '40');
INSERT INTO `sys_role_authority` VALUES ('15', '42');
INSERT INTO `sys_role_authority` VALUES ('15', '43');
INSERT INTO `sys_role_authority` VALUES ('15', '44');
INSERT INTO `sys_role_authority` VALUES ('15', '45');
INSERT INTO `sys_role_authority` VALUES ('15', '46');
INSERT INTO `sys_role_authority` VALUES ('15', '47');
INSERT INTO `sys_role_authority` VALUES ('15', '48');
INSERT INTO `sys_role_authority` VALUES ('15', '52');
INSERT INTO `sys_role_authority` VALUES ('15', '53');
INSERT INTO `sys_role_authority` VALUES ('15', '54');
INSERT INTO `sys_role_authority` VALUES ('15', '55');
INSERT INTO `sys_role_authority` VALUES ('15', '56');
INSERT INTO `sys_role_authority` VALUES ('15', '57');
INSERT INTO `sys_role_authority` VALUES ('15', '58');
INSERT INTO `sys_role_authority` VALUES ('15', '59');
INSERT INTO `sys_role_authority` VALUES ('15', '60');
INSERT INTO `sys_role_authority` VALUES ('15', '61');
INSERT INTO `sys_role_authority` VALUES ('15', '62');
INSERT INTO `sys_role_authority` VALUES ('15', '63');
INSERT INTO `sys_role_authority` VALUES ('15', '64');
INSERT INTO `sys_role_authority` VALUES ('15', '65');
INSERT INTO `sys_role_authority` VALUES ('15', '66');
INSERT INTO `sys_role_authority` VALUES ('15', '67');
INSERT INTO `sys_role_authority` VALUES ('15', '68');
INSERT INTO `sys_role_authority` VALUES ('15', '69');
INSERT INTO `sys_role_authority` VALUES ('15', '70');
INSERT INTO `sys_role_authority` VALUES ('15', '71');
INSERT INTO `sys_role_authority` VALUES ('15', '72');
INSERT INTO `sys_role_authority` VALUES ('15', '73');
INSERT INTO `sys_role_authority` VALUES ('15', '74');
INSERT INTO `sys_role_authority` VALUES ('15', '75');
INSERT INTO `sys_role_authority` VALUES ('15', '76');
INSERT INTO `sys_role_authority` VALUES ('15', '77');
INSERT INTO `sys_role_authority` VALUES ('15', '78');
INSERT INTO `sys_role_authority` VALUES ('15', '79');
INSERT INTO `sys_role_authority` VALUES ('15', '80');
INSERT INTO `sys_role_authority` VALUES ('15', '81');
INSERT INTO `sys_role_authority` VALUES ('15', '82');
INSERT INTO `sys_role_authority` VALUES ('15', '83');
INSERT INTO `sys_role_authority` VALUES ('15', '84');
INSERT INTO `sys_role_authority` VALUES ('15', '85');
INSERT INTO `sys_role_authority` VALUES ('15', '86');
INSERT INTO `sys_role_authority` VALUES ('15', '87');
INSERT INTO `sys_role_authority` VALUES ('15', '88');
INSERT INTO `sys_role_authority` VALUES ('15', '89');
INSERT INTO `sys_role_authority` VALUES ('15', '90');
INSERT INTO `sys_role_authority` VALUES ('15', '91');
INSERT INTO `sys_role_authority` VALUES ('15', '92');
INSERT INTO `sys_role_authority` VALUES ('15', '93');
INSERT INTO `sys_role_authority` VALUES ('15', '94');
INSERT INTO `sys_role_authority` VALUES ('15', '95');
INSERT INTO `sys_role_authority` VALUES ('15', '96');
INSERT INTO `sys_role_authority` VALUES ('15', '97');
INSERT INTO `sys_role_authority` VALUES ('15', '98');
INSERT INTO `sys_role_authority` VALUES ('15', '99');
INSERT INTO `sys_role_authority` VALUES ('15', '100');
INSERT INTO `sys_role_authority` VALUES ('15', '101');
INSERT INTO `sys_role_authority` VALUES ('15', '102');
INSERT INTO `sys_role_authority` VALUES ('15', '103');
INSERT INTO `sys_role_authority` VALUES ('15', '104');
INSERT INTO `sys_role_authority` VALUES ('15', '105');
INSERT INTO `sys_role_authority` VALUES ('15', '106');
INSERT INTO `sys_role_authority` VALUES ('15', '107');
INSERT INTO `sys_role_authority` VALUES ('15', '108');
INSERT INTO `sys_role_authority` VALUES ('15', '109');
INSERT INTO `sys_role_authority` VALUES ('15', '110');
INSERT INTO `sys_role_authority` VALUES ('15', '111');
INSERT INTO `sys_role_authority` VALUES ('15', '112');
INSERT INTO `sys_role_authority` VALUES ('15', '113');
INSERT INTO `sys_role_authority` VALUES ('15', '114');
INSERT INTO `sys_role_authority` VALUES ('15', '115');
INSERT INTO `sys_role_authority` VALUES ('15', '116');
INSERT INTO `sys_role_authority` VALUES ('15', '117');
INSERT INTO `sys_role_authority` VALUES ('15', '118');
INSERT INTO `sys_role_authority` VALUES ('15', '119');
INSERT INTO `sys_role_authority` VALUES ('15', '120');
INSERT INTO `sys_role_authority` VALUES ('15', '121');
INSERT INTO `sys_role_authority` VALUES ('15', '122');
INSERT INTO `sys_role_authority` VALUES ('15', '123');
INSERT INTO `sys_role_authority` VALUES ('15', '124');
INSERT INTO `sys_role_authority` VALUES ('15', '125');
INSERT INTO `sys_role_authority` VALUES ('15', '126');
INSERT INTO `sys_role_authority` VALUES ('15', '127');
INSERT INTO `sys_role_authority` VALUES ('15', '128');
INSERT INTO `sys_role_authority` VALUES ('15', '129');
INSERT INTO `sys_role_authority` VALUES ('15', '130');
INSERT INTO `sys_role_authority` VALUES ('15', '131');
INSERT INTO `sys_role_authority` VALUES ('15', '132');
INSERT INTO `sys_role_authority` VALUES ('15', '133');
INSERT INTO `sys_role_authority` VALUES ('15', '134');
INSERT INTO `sys_role_authority` VALUES ('15', '135');
INSERT INTO `sys_role_authority` VALUES ('15', '136');
INSERT INTO `sys_role_authority` VALUES ('15', '137');
INSERT INTO `sys_role_authority` VALUES ('15', '138');
INSERT INTO `sys_role_authority` VALUES ('15', '139');
INSERT INTO `sys_role_authority` VALUES ('15', '140');
INSERT INTO `sys_role_authority` VALUES ('15', '141');
INSERT INTO `sys_role_authority` VALUES ('15', '142');
INSERT INTO `sys_role_authority` VALUES ('15', '143');
INSERT INTO `sys_role_authority` VALUES ('15', '144');
INSERT INTO `sys_role_authority` VALUES ('15', '145');
INSERT INTO `sys_role_authority` VALUES ('15', '146');
INSERT INTO `sys_role_authority` VALUES ('15', '147');
INSERT INTO `sys_role_authority` VALUES ('15', '148');
INSERT INTO `sys_role_authority` VALUES ('15', '151');
INSERT INTO `sys_role_authority` VALUES ('15', '152');
INSERT INTO `sys_role_authority` VALUES ('15', '153');
INSERT INTO `sys_role_authority` VALUES ('15', '154');
INSERT INTO `sys_role_authority` VALUES ('15', '155');
INSERT INTO `sys_role_authority` VALUES ('15', '156');
INSERT INTO `sys_role_authority` VALUES ('15', '157');
INSERT INTO `sys_role_authority` VALUES ('15', '158');
INSERT INTO `sys_role_authority` VALUES ('16', '1');
INSERT INTO `sys_role_authority` VALUES ('16', '2');
INSERT INTO `sys_role_authority` VALUES ('16', '3');
INSERT INTO `sys_role_authority` VALUES ('16', '4');
INSERT INTO `sys_role_authority` VALUES ('16', '5');
INSERT INTO `sys_role_authority` VALUES ('16', '8');
INSERT INTO `sys_role_authority` VALUES ('16', '9');
INSERT INTO `sys_role_authority` VALUES ('16', '10');
INSERT INTO `sys_role_authority` VALUES ('16', '11');
INSERT INTO `sys_role_authority` VALUES ('16', '12');

-- ----------------------------
-- Table structure for sys_sign
-- ----------------------------
DROP TABLE IF EXISTS `sys_sign`;
CREATE TABLE `sys_sign` (
  `sign_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `integral` bigint(20) DEFAULT NULL,
  `count` int(10) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`sign_id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_sign
-- ----------------------------
INSERT INTO `sys_sign` VALUES ('57', '1', '1322', '1', '2021-11-17 15:59:06');

-- ----------------------------
-- Table structure for sys_tag
-- ----------------------------
DROP TABLE IF EXISTS `sys_tag`;
CREATE TABLE `sys_tag` (
  `tag_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `top` tinyint(4) DEFAULT '1' COMMENT '是否推荐 1不推荐 2推荐',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_tag
-- ----------------------------
INSERT INTO `sys_tag` VALUES ('1', '小狗', '1', '2020-11-05 00:27:18');
INSERT INTO `sys_tag` VALUES ('2', '恐龙', '1', '2020-11-05 00:27:46');
INSERT INTO `sys_tag` VALUES ('3', '鸿鹄', '1', '2020-11-05 00:28:16');
INSERT INTO `sys_tag` VALUES ('4', '白猪', '1', '2020-11-05 00:29:34');
INSERT INTO `sys_tag` VALUES ('5', '鹦鹉', '1', '2020-11-05 00:30:04');
INSERT INTO `sys_tag` VALUES ('6', '普虎', '1', '2020-11-05 00:30:24');
INSERT INTO `sys_tag` VALUES ('7', '黑鹤', '1', '2020-11-05 00:30:44');
INSERT INTO `sys_tag` VALUES ('8', '牛奇', '1', '2020-11-05 00:31:01');
INSERT INTO `sys_tag` VALUES ('12', '啊实打', '1', '2020-11-06 10:29:18');
INSERT INTO `sys_tag` VALUES ('14', '豆腐干地方', '1', '2020-11-08 02:19:32');
INSERT INTO `sys_tag` VALUES ('15', 'FSGHSDFGSD', '1', '2020-11-08 02:21:53');
INSERT INTO `sys_tag` VALUES ('16', 'sdfasd ', '1', '2020-11-08 23:16:21');
INSERT INTO `sys_tag` VALUES ('17', 'gh', '1', '2020-11-08 23:16:21');
INSERT INTO `sys_tag` VALUES ('18', 'ghjfg', '1', '2020-11-08 23:23:19');
INSERT INTO `sys_tag` VALUES ('19', 'sdafs', '1', '2020-11-08 23:28:31');
INSERT INTO `sys_tag` VALUES ('20', 'bvjg', '1', '2020-11-08 23:31:57');
INSERT INTO `sys_tag` VALUES ('21', 'dsfgsdf', '1', '2020-11-08 23:31:57');
INSERT INTO `sys_tag` VALUES ('22', 'fdgsdfg fdg', '1', '2020-12-08 16:54:47');
INSERT INTO `sys_tag` VALUES ('23', '艾什', '1', '2021-02-19 20:50:59');
INSERT INTO `sys_tag` VALUES ('24', '撒打发', '1', '2021-02-19 21:21:16');
INSERT INTO `sys_tag` VALUES ('25', '豆腐干岁的法国', '1', '2021-02-23 22:35:37');
INSERT INTO `sys_tag` VALUES ('26', '风格的和', '1', '2021-03-07 17:29:28');
INSERT INTO `sys_tag` VALUES ('27', '士大夫', '1', '2021-08-24 19:45:25');
INSERT INTO `sys_tag` VALUES ('28', '哇撒入发', '1', '2021-08-25 00:58:54');
INSERT INTO `sys_tag` VALUES ('29', '豆腐干大师傅给', '1', '2021-08-25 01:03:40');
INSERT INTO `sys_tag` VALUES ('30', '环球时报', '1', '2021-08-26 16:13:06');
INSERT INTO `sys_tag` VALUES ('31', '中国新闻', '1', '2021-08-26 16:14:50');
INSERT INTO `sys_tag` VALUES ('32', '超2000亿美元的公司', '1', '2021-08-26 16:18:34');
INSERT INTO `sys_tag` VALUES ('33', 'apex', '1', '2021-08-26 18:16:26');
INSERT INTO `sys_tag` VALUES ('34', '测试课程', '1', '2021-08-26 18:34:40');
INSERT INTO `sys_tag` VALUES ('35', '应用', '1', '2021-08-26 18:44:43');
INSERT INTO `sys_tag` VALUES ('36', 'Photoshop', '1', '2021-11-19 17:21:37');
INSERT INTO `sys_tag` VALUES ('37', 'procreate', '1', '2021-11-19 17:22:11');
INSERT INTO `sys_tag` VALUES ('38', 'snipaste', '1', '2021-11-19 17:22:43');
INSERT INTO `sys_tag` VALUES ('39', 'Krita', '1', '2021-11-19 17:24:21');
INSERT INTO `sys_tag` VALUES ('40', 'artstudio', '1', '2021-11-19 17:25:04');
INSERT INTO `sys_tag` VALUES ('41', 'LIVE2D', '1', '2021-11-19 17:26:10');
INSERT INTO `sys_tag` VALUES ('42', '广告', '1', '2021-11-19 17:41:34');
INSERT INTO `sys_tag` VALUES ('43', '模板', '1', '2021-11-19 17:43:23');
INSERT INTO `sys_tag` VALUES ('44', 'AE', '1', '2021-11-19 17:44:38');
INSERT INTO `sys_tag` VALUES ('45', 'MV', '1', '2021-11-19 17:46:30');
INSERT INTO `sys_tag` VALUES ('46', '采样音色', '1', '2021-11-19 17:52:52');
INSERT INTO `sys_tag` VALUES ('47', 'beat', '1', '2021-11-19 17:54:30');
INSERT INTO `sys_tag` VALUES ('48', '设计思路', '1', '2021-11-19 21:03:40');
INSERT INTO `sys_tag` VALUES ('49', 'Kontakt', '1', '2021-11-19 21:07:54');
INSERT INTO `sys_tag` VALUES ('50', 'Android 12', '1', '2021-11-19 21:30:21');
INSERT INTO `sys_tag` VALUES ('51', 'Fiber安装教程', '1', '2021-12-15 22:33:31');
INSERT INTO `sys_tag` VALUES ('52', 'sadf', '0', '2022-01-11 22:27:01');

-- ----------------------------
-- Table structure for sys_tag_related
-- ----------------------------
DROP TABLE IF EXISTS `sys_tag_related`;
CREATE TABLE `sys_tag_related` (
  `tag_id` bigint(20) NOT NULL,
  `related_id` bigint(20) NOT NULL COMMENT '关系id',
  `module` varchar(50) DEFAULT NULL COMMENT '模块'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_tag_related
-- ----------------------------
INSERT INTO `sys_tag_related` VALUES ('3', '1', 'resource');
INSERT INTO `sys_tag_related` VALUES ('4', '1', 'resource');
INSERT INTO `sys_tag_related` VALUES ('36', '4', 'resource');
INSERT INTO `sys_tag_related` VALUES ('38', '3', 'resource');
INSERT INTO `sys_tag_related` VALUES ('37', '2', 'resource');
INSERT INTO `sys_tag_related` VALUES ('39', '5', 'resource');
INSERT INTO `sys_tag_related` VALUES ('40', '6', 'resource');
INSERT INTO `sys_tag_related` VALUES ('41', '7', 'resource');
INSERT INTO `sys_tag_related` VALUES ('42', '2', 'video');
INSERT INTO `sys_tag_related` VALUES ('43', '3', 'video');
INSERT INTO `sys_tag_related` VALUES ('44', '4', 'video');
INSERT INTO `sys_tag_related` VALUES ('45', '5', 'video');
INSERT INTO `sys_tag_related` VALUES ('46', '3', 'audio');
INSERT INTO `sys_tag_related` VALUES ('47', '2', 'audio');
INSERT INTO `sys_tag_related` VALUES ('15', '4', 'audio');
INSERT INTO `sys_tag_related` VALUES ('17', '5', 'audio');
INSERT INTO `sys_tag_related` VALUES ('25', '6', 'audio');
INSERT INTO `sys_tag_related` VALUES ('17', '7', 'audio');
INSERT INTO `sys_tag_related` VALUES ('3', '2', 'edu');
INSERT INTO `sys_tag_related` VALUES ('6', '2', 'edu');
INSERT INTO `sys_tag_related` VALUES ('48', '3', 'edu');
INSERT INTO `sys_tag_related` VALUES ('49', '4', 'edu');
INSERT INTO `sys_tag_related` VALUES ('25', '5', 'edu');
INSERT INTO `sys_tag_related` VALUES ('32', '2', 'article');
INSERT INTO `sys_tag_related` VALUES ('50', '3', 'article');
INSERT INTO `sys_tag_related` VALUES ('31', '1', 'article');
INSERT INTO `sys_tag_related` VALUES ('25', '4', 'article');
INSERT INTO `sys_tag_related` VALUES ('51', '5', 'article');

-- ----------------------------
-- Table structure for sys_topic
-- ----------------------------
DROP TABLE IF EXISTS `sys_topic`;
CREATE TABLE `sys_topic` (
  `topic_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '发布用户',
  `group_id` bigint(20) DEFAULT NULL,
  `related_id` bigint(20) DEFAULT NULL COMMENT '关联内容的id',
  `module` varchar(50) DEFAULT NULL COMMENT '所属模块',
  `type` tinyint(4) NOT NULL COMMENT '类型 1帖子，2问答',
  `title` varchar(150) DEFAULT NULL COMMENT '话题',
  `files` text COMMENT '文件链接',
  `views` bigint(20) DEFAULT NULL COMMENT '查看',
  `likes` bigint(20) DEFAULT NULL,
  `hots` bigint(20) DEFAULT '0',
  `is_top` tinyint(4) DEFAULT NULL COMMENT '是否置顶 1 不置顶 2置顶',
  `price` decimal(10,2) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`topic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_topic
-- ----------------------------
INSERT INTO `sys_topic` VALUES ('1', '3', '2', null, null, '1', '这是一条测试动态', '', '0', '0', '0', '2', '0.00', '1', '2021-05-11 06:08:58', '2021-05-11 06:08:58', null, '通过');
INSERT INTO `sys_topic` VALUES ('2', '3', '2', null, null, '1', '测试话题看到法国佳士得风格的法国第三方??', '[\"http://localhost:8199/public/uploads/2021-07-13/ccrajvviahj4dv0p5d.png\"]', '4', '0', '0', '1', '0.00', '2', '2021-07-13 01:08:18', '2021-07-13 01:08:18', null, '');
INSERT INTO `sys_topic` VALUES ('3', '1', '2', null, null, '1', '撒旦撒发射点发射点发噶啥的噶啥的??', '[\"http://localhost:8199/public/uploads/2021-11-19/cftnqt0817ekqffczc.jpg\"]', '43', '0', '0', '1', '0.00', '2', '2021-07-21 00:34:48', '2021-07-21 00:34:48', null, '');
INSERT INTO `sys_topic` VALUES ('4', '1', '2', null, null, '1', '的风沙的方式规划建设大概阿斯顿飞过是对方进攻', '[\"http://localhost:8199/public/uploads/2021-07-13/ccrajvviahj4dv0p5d.png\"]', '2', '0', '0', '1', '0.00', '1', '2021-08-24 01:04:08', '2021-08-24 01:04:08', null, '');
INSERT INTO `sys_topic` VALUES ('5', '1', '2', null, null, '1', '规范化岁的法国还是大发噶的风格', '[\"http://localhost:8199/public/uploads/2021-07-13/ccrajvviahj4dv0p5d.png\"]', '3', '0', '0', '1', '0.00', '2', '2021-08-24 01:05:20', '2021-08-24 01:05:20', null, '');
INSERT INTO `sys_topic` VALUES ('6', '1', '2', null, null, '1', '风格化风格化是风格豆腐干', '', '0', '0', '0', '1', '0.00', '2', '2021-10-28 01:16:35', '2021-10-28 01:16:35', null, '');
INSERT INTO `sys_topic` VALUES ('7', '1', '2', null, null, '1', 'fdxgsdfgzdfgsdfg', '', '2', '0', '0', '1', '0.00', '2', '2021-10-28 23:04:19', '2021-10-28 23:04:19', null, '');
INSERT INTO `sys_topic` VALUES ('8', '1', '2', null, null, '1', 'dfgsdfgsdfgsfdghdfghdfgh', '', '1', '0', '0', '1', '0.00', '2', '2021-10-28 23:09:50', '2021-10-28 23:09:50', null, '');
INSERT INTO `sys_topic` VALUES ('9', '1', '2', null, null, '1', 'fghjfghjgfhjfghjgfhjg', '', '11', '0', '0', '1', '0.00', '2', '2021-10-28 23:10:30', '2021-10-28 23:10:30', null, '');
INSERT INTO `sys_topic` VALUES ('14', '1', '2', '1', 'article', '2', '给v换成风格和非递归', '', '27', '0', '0', '1', '0.00', '2', '2021-11-11 01:04:43', '2021-11-11 01:04:43', null, '');
INSERT INTO `sys_topic` VALUES ('15', '1', '2', '0', '', '1', '风格还是的风格和岁的法国很多想法是', '', '0', '0', '0', '1', '0.00', '2', '2021-12-01 17:19:28', '2021-12-01 17:19:28', null, '');
INSERT INTO `sys_topic` VALUES ('16', '1', '2', '0', '', '1', '风格还是的风格和岁的法国很多想法是', '[\"http://localhost:8199/public/uploads/2021-12-01/cg3uqqylnfccaipjrs.jpg\"]', '5', '0', '0', '1', '0.00', '2', '2021-12-01 17:19:51', '2021-12-01 17:19:51', null, '');
INSERT INTO `sys_topic` VALUES ('17', '1', '2', '0', '', '1', '非常v宝宝相册v表现出v表现出v表现出', '', '0', '0', '0', '1', '0.00', '2', '2021-12-01 17:20:33', '2021-12-01 17:20:33', null, '');
INSERT INTO `sys_topic` VALUES ('18', '1', '2', '0', '', '1', '大发噶大发噶大发噶当时法国', '', '0', '0', '0', '1', '0.00', '2', '2021-12-01 17:21:35', '2021-12-01 17:21:35', null, '');
INSERT INTO `sys_topic` VALUES ('19', '1', '2', '0', '', '1', 'u急口令回家看了回家看了回家看了', '', '0', '0', '0', '1', '0.00', '2', '2021-12-01 17:22:17', '2021-12-01 17:22:17', null, '');
INSERT INTO `sys_topic` VALUES ('20', '1', '2', '0', '', '1', '岁的法国撒的分公司的还是的风格和', '', '3', '1', '0', '1', '0.00', '2', '2021-12-01 17:23:42', '2021-12-01 17:23:42', null, '');
INSERT INTO `sys_topic` VALUES ('21', '1', '2', '0', '', '1', '飞过的痕迹帝国海军的风格和的风格和', '', '6', '0', '0', '1', '0.00', '2', '2021-12-01 17:24:55', '2021-12-01 17:24:55', null, '');
INSERT INTO `sys_topic` VALUES ('22', '1', '3', '0', '', '1', '的说法伽撒大噶的风格岁的法国', '', '0', '0', '0', '1', '0.00', '2', '2022-03-12 22:58:30', '2022-03-12 22:58:30', null, '');
INSERT INTO `sys_topic` VALUES ('23', '1', '2', '0', '', '1', '士大夫嘎斯的发嘎达嘎达发给第三方', '', '0', '0', '0', '1', '0.00', '2', '2022-03-12 23:03:35', '2022-03-12 23:03:35', null, '');
INSERT INTO `sys_topic` VALUES ('24', '1', '2', '0', '', '1', '点三公分SDF手动阀手动阀手动阀手动阀', '', '0', '0', '0', '1', '0.00', '2', '2022-03-12 23:04:07', '2022-03-12 23:04:07', null, '');
INSERT INTO `sys_topic` VALUES ('25', '1', '2', '0', '', '1', '的说法伽师的噶SDF士大夫', '', '1', '0', '0', '1', '0.00', '2', '2022-03-12 23:04:47', '2022-03-12 23:04:47', null, '');
INSERT INTO `sys_topic` VALUES ('26', '1', '3', '0', '', '1', '首都曼谷哈吉斯看到噶啥的 撒打发', '', '0', '0', '0', '1', '0.00', '2', '2022-03-14 16:23:42', '2022-03-14 16:23:42', null, '');
INSERT INTO `sys_topic` VALUES ('27', '1', '2', '0', '', '2', '符合公司法国还是发给很舒服当时法国', '', '4', '0', '0', '1', '0.00', '2', '2022-03-14 16:24:40', '2022-03-14 16:24:40', null, '');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nick_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(4) DEFAULT '3' COMMENT '用户性别（1男 2女 3未知）',
  `avatar` varchar(255) DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `salt` char(10) DEFAULT NULL COMMENT '密码盐',
  `cover` varchar(255) DEFAULT NULL,
  `follows` bigint(20) DEFAULT NULL,
  `fans` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `grade` bigint(20) DEFAULT NULL,
  `vip` bigint(20) DEFAULT NULL,
  `balance` decimal(10,2) DEFAULT NULL COMMENT '余额',
  `integral` bigint(20) DEFAULT NULL COMMENT '积分',
  `description` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0' COMMENT '帐号状态（1停用,2正常）',
  `login_ip` varchar(50) DEFAULT '' COMMENT '最后登陆IP',
  `login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `delete_time` datetime DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'EForinaj', 'fiber@admin.com', '', '1', 'http://localhost:8199/public/uploads/2021-11-19/cftsrl7r9nnngyv4ku.jpg', 'f3565ae9276f3b7244230c1bcbb4ff4c', 'ePJsln', 'http://localhost:8199/public/uploads/2021-11-19/cftsrmz4obqsgk9ndy.jpg', '1', '1', '11', '3', null, '538.81', '1200', '干撒大发噶丹是', '2', '127.0.0.1', '2022-03-30 19:09:23', '2021-05-10 05:50:02', '2021-08-26 18:39:21', '', null);
INSERT INTO `sys_user` VALUES ('3', '新用户248752', 'test@admin.com', '', '1', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', 'f3565ae9276f3b7244230c1bcbb4ff4c', 'ePJsln', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', '13', '13', '126', '2', '1', '76.91', '450', '风格和对方水果和法国的还是的风格和是的', '2', '[::1]', '2021-11-16 17:52:28', '2021-05-10 05:58:47', '2021-08-26 11:41:24', '', null);
INSERT INTO `sys_user` VALUES ('4', '新用户869377', 'mushokumunou@gmail.com', '', '3', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', '75965be72c6c21f7345789ce547c2ec9', 'gCiJ6S', 'http://localhost:8199/public/uploads/2021-07-05/ccl6yfkhed74e6dztr.png', '0', '0', '0', '1', '0', '3.23', '123', '', '2', '178.132.6.37', '2021-05-11 10:40:24', '2021-05-11 10:40:14', '2021-06-10 14:37:44', '', null);
INSERT INTO `sys_user` VALUES ('5', '新用户149587', '356866114@qq.com', '', '3', 'http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png', '4429df92f33044f8238892f5a8bf0b6e', 'H3snbI', 'http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png', '0', '0', '0', '1', '0', '0.00', '300', '', '2', '[::1]', '2021-11-28 20:28:15', '2021-11-28 20:28:15', '2021-11-28 20:28:15', '', null);
INSERT INTO `sys_user` VALUES ('6', '新用户702149', 'aganlaizui@qq.com', '', '3', 'http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png', '3fc1d1930238201faa64e01dc3bb1422', 'Uv0N96', 'http://localhost:8199/public/uploads/2021-06-30/ccgxrrsx6enov4utih.png', '0', '0', '0', '1', '0', '0.00', '300', '', '2', '60.206.118.29', '2022-01-03 23:48:50', '2022-01-03 23:47:32', '2022-01-03 23:47:32', '', null);

-- ----------------------------
-- Table structure for sys_user_favorite
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_favorite`;
CREATE TABLE `sys_user_favorite` (
  `user_id` bigint(20) DEFAULT NULL,
  `favorite_id` bigint(20) DEFAULT NULL,
  `module` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_favorite
-- ----------------------------
INSERT INTO `sys_user_favorite` VALUES ('1', '2', 'question');
INSERT INTO `sys_user_favorite` VALUES ('1', '3', 'article');

-- ----------------------------
-- Table structure for sys_user_follow
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_follow`;
CREATE TABLE `sys_user_follow` (
  `user_id` bigint(20) DEFAULT NULL,
  `follow_id` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_follow
-- ----------------------------
INSERT INTO `sys_user_follow` VALUES ('1', '3');
INSERT INTO `sys_user_follow` VALUES ('3', '1');

-- ----------------------------
-- Table structure for sys_user_join_edu
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_join_edu`;
CREATE TABLE `sys_user_join_edu` (
  `mode` tinyint(4) DEFAULT NULL COMMENT '联系方式1微信，2QQ，3手机',
  `number` varchar(50) DEFAULT NULL,
  `edu_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_join_edu
-- ----------------------------
INSERT INTO `sys_user_join_edu` VALUES ('1', 'DFSGASD', '2', '1', 'dfgasdf', '2021-07-16 23:59:10');

-- ----------------------------
-- Table structure for sys_user_join_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_join_group`;
CREATE TABLE `sys_user_join_group` (
  `group_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_join_group
-- ----------------------------
INSERT INTO `sys_user_join_group` VALUES ('2', '1');

-- ----------------------------
-- Table structure for sys_user_like
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_like`;
CREATE TABLE `sys_user_like` (
  `user_id` bigint(20) NOT NULL,
  `related_id` bigint(20) NOT NULL,
  `module` varchar(50) NOT NULL,
  PRIMARY KEY (`user_id`,`related_id`,`module`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_like
-- ----------------------------
INSERT INTO `sys_user_like` VALUES ('1', '1', 'comment');
INSERT INTO `sys_user_like` VALUES ('1', '2', 'article');
INSERT INTO `sys_user_like` VALUES ('1', '3', 'article');
INSERT INTO `sys_user_like` VALUES ('1', '4', 'answer');
INSERT INTO `sys_user_like` VALUES ('1', '20', 'topic');
INSERT INTO `sys_user_like` VALUES ('1', '49', 'comment');
INSERT INTO `sys_user_like` VALUES ('1', '50', 'comment');
INSERT INTO `sys_user_like` VALUES ('1', '53', 'comment');
INSERT INTO `sys_user_like` VALUES ('1', '54', 'comment');
INSERT INTO `sys_user_like` VALUES ('3', '1', 'article');
INSERT INTO `sys_user_like` VALUES ('3', '1', 'comment');
INSERT INTO `sys_user_like` VALUES ('3', '3', 'article');
INSERT INTO `sys_user_like` VALUES ('3', '5', 'question');
INSERT INTO `sys_user_like` VALUES ('3', '22', 'answer');

-- ----------------------------
-- Table structure for sys_user_vip
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_vip`;
CREATE TABLE `sys_user_vip` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `vip_id` bigint(20) NOT NULL COMMENT '角色ID',
  `start_time` datetime DEFAULT NULL,
  `finish_time` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`,`vip_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_vip
-- ----------------------------
INSERT INTO `sys_user_vip` VALUES ('3', '1', '2021-08-26 11:39:59', '2021-09-26 11:39:59');

-- ----------------------------
-- Table structure for sys_verify
-- ----------------------------
DROP TABLE IF EXISTS `sys_verify`;
CREATE TABLE `sys_verify` (
  `verify_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `code` varchar(50) DEFAULT NULL COMMENT '身份证号码',
  `encryption` varchar(255) DEFAULT NULL COMMENT '加密后的认证信息',
  `mode` tinyint(4) DEFAULT NULL COMMENT '联系方式 1 qq, 2微信',
  `number` varchar(50) DEFAULT NULL COMMENT '联系号码',
  `create_time` datetime DEFAULT NULL COMMENT '认证时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态（1待审，2成功，3拒绝）',
  PRIMARY KEY (`verify_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_verify
-- ----------------------------
INSERT INTO `sys_verify` VALUES ('3', '1', '荣*来', '455******654', '1a23d963924fc704f0c3132d871222e0', '1', '6565156', '2021-07-10 17:14:26', '2021-07-10 17:14:26', '', '2');

-- ----------------------------
-- Table structure for sys_video
-- ----------------------------
DROP TABLE IF EXISTS `sys_video`;
CREATE TABLE `sys_video` (
  `video_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `cate_id` bigint(20) DEFAULT NULL COMMENT '分类',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `link` varchar(255) DEFAULT NULL COMMENT '视频地址',
  `hots` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL COMMENT '点赞数',
  `favorites` bigint(20) unsigned DEFAULT NULL COMMENT '收藏',
  `views` bigint(20) DEFAULT NULL COMMENT '播放量',
  `has_down` tinyint(4) DEFAULT NULL COMMENT '是否有下载1没有，2有',
  `down_mode` tinyint(4) DEFAULT NULL COMMENT '下载权限 0公开下载，1付费下载，2评论下载，3登录下载',
  `price` decimal(10,2) DEFAULT NULL,
  `down_url` text,
  `purpose` text,
  `attribute` text,
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `other_link` varchar(255) DEFAULT NULL COMMENT '第三方地址',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿',
  `delete_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_video
-- ----------------------------
INSERT INTO `sys_video` VALUES ('2', '1', '6', '40+个整洁动态文本文字标题呼出线动画AE视频模板素材', 'http://localhost:8199/public/uploads/2021-11-19/cftnonwtwzmjcc0hwi.jpg', 'http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', '1231', '12123', '123', '136', '2', '1', '23.00', '[{\"key\":\"fibercms.com\",\"title\":\"百度网盘\",\"val\":\"sdf\"}]', '[{\"key\":\"商业使用\",\"val\":\"2\"}]', '[{\"key\":\"xxx\",\"val\":\"xxxx\"}]', '不拿狗子 拿什么希尔', '', '2', null, '2021-07-12 17:48:40', '2021-11-19 17:41:34', '');
INSERT INTO `sys_video` VALUES ('3', '1', '6', '抽象唯美霓虹粒子线条文字标题开场片头AE视频模板素材 Particle Titles', 'http://localhost:8199/public/uploads/2021-11-19/cftnpvc6xash3icb7t.jpg', 'http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', '123', '123', '1231', '156', '1', '0', '0.00', '[]', '[]', '[]', '豆腐干士大夫敢死队风格是豆腐干豆腐干', '', '2', null, '2021-11-19 17:43:23', '2021-11-19 17:43:23', '');
INSERT INTO `sys_video` VALUES ('4', '1', '6', '20组可爱卡通文本文字标题气泡动画AE视频模板素材 Comic Titles After Effects', 'http://localhost:8199/public/uploads/2021-11-19/cftnqt0817ekqffczc.jpg', 'http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', '2131', '3123', '2312', '123127', '1', '0', '0.00', '[]', '[]', '[]', '卡通、时尚、多彩的文字动画模板。极其易于使用、精心分组和组织。调整颜色，输入文本，然后享受结果。', '', '2', null, '2021-11-19 17:44:38', '2021-11-19 17:44:38', '');
INSERT INTO `sys_video` VALUES ('5', '1', '6', '3D标题-照片标题动画图文展示视频AE模板', 'http://localhost:8199/public/uploads/2021-11-19/cftnsb3a88t58v3axf.png', 'http://localhost:8199/public/uploads/2021-11-19/cftnoto7hqia1j1rct.mp4', '21312', '312312', '31231', '23126', '1', '0', '0.00', '[]', '[]', '[]', '3D标题-照片标题是由envato上传素材并提供无水印可编辑AE模版源文件下载服务，视频分辨率为：1920x1080，喜欢当前栏目包装文字排版AE模板的就抓紧下载吧！\n\n', '', '2', null, '2021-11-19 17:46:30', '2021-11-19 17:46:30', '');

-- ----------------------------
-- Table structure for sys_vip
-- ----------------------------
DROP TABLE IF EXISTS `sys_vip`;
CREATE TABLE `sys_vip` (
  `vip_id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(30) DEFAULT NULL COMMENT '会员标题',
  `icon` varchar(255) DEFAULT NULL,
  `day` int(4) DEFAULT NULL COMMENT '到期时间',
  `price` decimal(10,2) DEFAULT NULL COMMENT '开通价格',
  `discount` decimal(10,2) DEFAULT NULL COMMENT '折扣',
  `color` varchar(30) DEFAULT NULL COMMENT '自定义颜色',
  PRIMARY KEY (`vip_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_vip
-- ----------------------------
INSERT INTO `sys_vip` VALUES ('1', '月会员', 'http://localhost:8199/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', '31', '20.00', '0.00', '');
INSERT INTO `sys_vip` VALUES ('2', '年会员', 'http://localhost:8199/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', '365', '2000.00', '0.00', '');
INSERT INTO `sys_vip` VALUES ('3', '永久会员', 'http://localhost:8199/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', '0', '3423.00', '0.00', '');
INSERT INTO `sys_vip` VALUES ('4', '测试会员第三方', 'http://localhost:8199/public/uploads/2021-11-19/cftswh0hteh6uoswhh.png', '50', '324.00', '0.00', '0');
