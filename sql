-- 用户表
CREATE TABLE user (
  id INT PRIMARY KEY AUTO_INCREMENT, -- 用户ID
  username VARCHAR(20) NOT NULL UNIQUE, -- 用户名
  password VARCHAR(32) NOT NULL, -- 密码
  nickname VARCHAR(20) NOT NULL, -- 昵称
  avatar VARCHAR(100), -- 头像
  gender CHAR(1), -- 性别
  birthday DATE -- 生日
);

-- 视频表
CREATE TABLE video (
  id INT PRIMARY KEY AUTO_INCREMENT, -- 视频ID
  title VARCHAR(50) NOT NULL, -- 标题
  description VARCHAR(200), -- 描述
  cover VARCHAR(100), -- 封面
  duration INT NOT NULL, -- 时长（秒）
  play_count INT DEFAULT 0, -- 播放量
  like_count INT DEFAULT 0, -- 点赞数
);

-- 音频表
CREATE TABLE audio (
  id INT PRIMARY KEY AUTO_INCREMENT, -- 音频ID
  name VARCHAR(50) NOT NULL, -- 名称
  author VARCHAR(20), --作者 
   duration INT NOT NULL--时长（秒）
);

--评论表 
CREATE TABLE comment (
   id INT PRIMARY KEY AUTO_INCREMENT ,--评论 ID 
   user_id INT NOT NULL ,--用户 ID 
   video_id INT NOT NULL ,--视频 ID 
   content VARCHAR (200)NOT NULL ,--内容 
   create_time DATETIME DEFAULT NOW ()--创建时间 
);

--点赞 表 
CREATE TABLE like (
   id INT PRIMARY KEY AUTO_INCREMENT ,--点赞 ID 
   user_id INT NOT NULL ,--用户 ID 
   target_id INT NOT NULL ,--目标 ID （视频 ID 或评论 ID）
);