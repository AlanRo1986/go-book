-- phpMyAdmin SQL Dump
-- version phpStudy 2014
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2019 �?05 �?07 �?01:00
-- 服务器版本: 5.5.53
-- PHP 版本: 5.5.38

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `lx_book`
--

-- --------------------------------------------------------

--
-- 表的结构 `lx_article`
--

CREATE TABLE IF NOT EXISTS `lx_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userId` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` longtext,
  `createTime` int(11) NOT NULL,
  `readCount` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `userId` (`userId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=9 ;

--
-- 转存表中的数据 `lx_article`
--

INSERT INTO `lx_article` (`id`, `userId`, `title`, `content`, `createTime`, `readCount`) VALUES
(5, 1, '服务器数据迁移', 'asdgas', 1557046670, 10),
(8, 1, '进入我岛礁12海里', '环球网5月6日消息，美国军舰6日再闯南沙海域。路透社援引美国军方说法称，美国两艘导弹驱逐舰周一(6日)进入中国南沙岛礁12海里范围内航行。\n\n路透社刚刚发布的一则“独家报道”称，美国军方发言人表示，美国“普雷布尔”号和“钟云”号导弹驱逐舰6日进入中国南沙南薰礁与赤瓜礁12海里内水域。这则“独家报道”还援引援引美国第七舰队发言人克莱·多斯指挥官的话，老调重弹地宣称，“无害通过”是“按照国际法的规定保护水路通道”。\n\n近期，美国军舰多次擅闯南海，被解放军警告驱离。今年2月11日，美国“斯普鲁恩斯”号和“普雷贝尔”号军舰，未经中国政府允许，擅自进入中国南沙群岛仁爱礁和美济礁邻近海域。当时，中国海军依法对美舰进行了识别查证，并且予以了警告驱离。\n\n外交部发言人华春莹当天对此表示，美方军舰的有关行为侵犯中国的主权，破坏有关海域的和平、安全和良好秩序。中方对此表示强烈的不满和坚决的反对。中国对南沙群岛及其附近海域拥有无可争辩的主权，中方一向尊重和维护各国依据国际法在南海享有的航行与飞越自由。但是坚决反对任何国家假借航行与飞越自由之名，损害沿岸国主权和安全。\n\n中国国防部此前也多次表态称，中国对南海诸岛及其附近海域拥有无可争辩的主权。中国军队将坚定履行防卫职责，继续采取一切必要措施，坚决捍卫国家主权安全，坚定维护地区和平稳定。\n\n（原题为《挑衅！美两艘导弹驱逐舰擅闯南沙，进入我岛礁12海里》）', 1557190794, 5);

-- --------------------------------------------------------

--
-- 表的结构 `lx_token`
--

CREATE TABLE IF NOT EXISTS `lx_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(32) NOT NULL,
  `status` int(1) NOT NULL DEFAULT '1',
  `userId` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `userId` (`userId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=12 ;

--
-- 转存表中的数据 `lx_token`
--

INSERT INTO `lx_token` (`id`, `token`, `status`, `userId`) VALUES
(1, '150841df15b542dfd153eafdf07d15c5', 0, 1),
(2, '4dcd32046636822c5462c29b5b8e1edd', 0, 1),
(3, '27e2ae56bdd11ad2abe13bd850bbe8cc', 0, 1),
(4, '0d6dc691ec368072885f4ad12e37fcc2', 0, 1),
(5, 'e155776f2a4e6bbf3f465c74f446a1ed', 0, 1),
(6, '04aabf518c747508207c4090ec70b7a3', 0, 1),
(7, '99fef40c42a9786a682d9675a9b54c5d', 0, 1),
(8, '60469943526717da2e45dd87570cc2e5', 0, 1),
(9, 'abc31f7c127504196640036b1f5099f6', 0, 1),
(10, '0e7f6839d07fd5c3de46553c3c37ca54', 0, 1),
(11, '9e26a53904d57f641de44315228f09b3', 1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `lx_user`
--

CREATE TABLE IF NOT EXISTS `lx_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `userpwd` varchar(32) NOT NULL,
  `createTime` int(11) NOT NULL,
  `loginIp` varchar(15) NOT NULL DEFAULT '""',
  `loginTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=6 ;

--
-- 转存表中的数据 `lx_user`
--

INSERT INTO `lx_user` (`id`, `username`, `userpwd`, `createTime`, `loginIp`, `loginTime`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1557029745, '127::0', 1557042745),
(2, 'admin1', 'fe3a0eada6dae0fbcf04f4968076b3e1', 1557042003, '', 0),
(3, 'admin0', '5259ee4a034fdeddd1b65be92debe731', 1557042192, '', 0),
(4, 'admin3', '5259ee4a034fdeddd1b65be92debe731', 1557042200, '', 0),
(5, 'admin4', '5259ee4a034fdeddd1b65be92debe731', 1557042627, '""', 0);

--
-- 限制导出的表
--

--
-- 限制表 `lx_article`
--
ALTER TABLE `lx_article`
  ADD CONSTRAINT `lx_article_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `lx_user` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;

--
-- 限制表 `lx_token`
--
ALTER TABLE `lx_token`
  ADD CONSTRAINT `lx_token_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `lx_user` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
