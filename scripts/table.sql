-- Active: 1738724774868@@172.16.6.69@3306@icu
-- auto-generated definition
# 用户表
create table icu.users (
`id` int auto_increment comment 'Primary Key' primary key,
`name` varchar(255) null comment 'Name',
`avatar` varchar(255) null comment 'avatar',
`email` varchar(255) null comment 'Email',
`username` varchar(255) null comment 'Username',
`signature` varchar(255) null comment 'signature',
`password` varchar(255) null comment 'Password',
`created_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'user';
#文章
create table icu.articles (
    `id` int auto_increment comment 'Primary Key' primary key,
    `title` varchar(255) null comment 'title',
    `content` TEXT null comment 'content',
    `user_id` varchar(255) null comment 'userId',
`created_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'article';

#标签表
create table icu.tags (
    `id` int auto_increment comment 'Primary Key' primary key,
    `tag` varchar(255) null comment 'tag',
    `user_id` varchar(255) null comment 'userId',
`created_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'tags';

#标签文章多对多关联表
create table icu.article_tags (
    `id` int auto_increment comment 'Primary Key' primary key,
    `article_id` int null comment 'articleId',
    `tag_id` int null comment 'tagId',
`created_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time',
    `deleted_at` datetime DEFAULT NULL null comment 'deleted_at',
    `updated_at` datetime DEFAULT NULL null comment 'updated_at'
) comment 'article_tags';

CREATE TABLE comments (
    `id` int auto_increment comment 'Primary Key' primary key,
    `article_id` BIGINT NOT NULL COMMENT 'Article ID',
    `user_id` BIGINT NOT NULL COMMENT 'User ID',
    `comment` TEXT NOT NULL,
    `parent_id` BIGINT DEFAULT NULL,
    `likes_count` INT DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'Create Time',
    `deleted_at` datetime DEFAULT NULL null comment 'deleted_at',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment 'updated_at'
);
ALTER TABLE icu.articles ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.tags ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.users ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.articles ADD COLUMN updated_at DATETIME NULL;

ALTER TABLE icu.tags ADD COLUMN updated_at DATETIME NULL;

ALTER TABLE icu.users ADD COLUMN updated_at DATETIME NULL;

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        1,
        1,
        1,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        2,
        1,
        6,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        3,
        2,
        2,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        4,
        2,
        8,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        5,
        3,
        3,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        6,
        3,
        7,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        7,
        4,
        4,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        8,
        4,
        9,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        9,
        5,
        5,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        10,
        5,
        9,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        11,
        6,
        6,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        12,
        6,
        8,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        13,
        7,
        2,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.article_tags (
        id,
        article_id,
        tag_id,
        create_at,
        deleted_at,
        updated_at
    )
VALUES (
        14,
        7,
        3,
        '2025-04-09 14:52:04',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        1,
        '2024年美国对华关税调整深度分析',
        '美国政府于2024年5月宣布对价值1800亿美元的中国商品加征关税，涉及电动汽车、锂电池及关键矿产，税率从7.5%提升至25%。此举直接影响特斯拉上海工厂对美出口计划，导致其股价单日下跌4.2%。中国商务部回应将采取反制措施，包括对美农产品加征15%报复性关税。专家预测，全球供应链可能向东南亚转移，越南和墨西哥的制造业投资同比增长23%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        2,
        'GPT-5与多模态AI的突破性进展',
        'OpenAI于2024年3月发布的GPT-5模型，参数规模达到1.8万亿，推理能力逼近人类专家水平。其多模态处理能力支持同时解析文本、图像与音频输入，在医疗诊断测试中，对X光片的异常检测准确率达98.7%。微软已将其集成到Azure云服务，每小时推理成本降低至0.12美元。伦理争议随之而来：欧盟AI法案要求所有生成内容必须添加数字水印。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        3,
        '劳动者必读：10步完成有效劳动仲裁',
        '根据《劳动争议调解仲裁法》，劳动者需在权益受损1年内提交仲裁申请。关键步骤包括：1）收集劳动合同、考勤记录；2）向企业注册地仲裁委提交申请书（模板可在人社部官网下载）；3）准备银行流水证明欠薪事实。2023年数据显示，北京地区仲裁胜诉率68%，平均赔偿金额为3.2个月工资。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        4,
        '拯救发际线：程序员防脱发全攻略',
        '长期熬夜与压力导致程序员脱发率高达43%（《2024程序员健康白皮书》）。建议：1）每日补充5mg生物素+1000IU维生素D；2）使用含咖啡因成分的防脱洗发液（如Alpecin C1）；3）每周进行2次头皮按摩（促进毛囊血液循环）。MIT研究显示，采用间歇式工作法（每编码45分钟休息15分钟）可降低皮质醇水平27%。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        5,
        '高尿酸血症饮食管理：从480μmol/L降到360的实战经验',
        '男性尿酸正常值应＜420μmol/L。实测发现，每日摄入嘌呤需控制在300mg以下：避免动物内脏（猪肝嘌呤含量275mg/100g），选择低脂乳制品。推荐食谱：早餐燕麦粥（嘌呤15mg）+ 午餐清蒸鲈鱼（82mg）+ 晚餐凉拌秋葵（60mg）。配合每日饮水2.5L，可提升尿酸排泄效率40%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        6,
        '2024全球芯片战争：美国对华技术封锁新规',
        '美国商务部5月更新《出口管制条例》，限制14nm以下制程设备对华出口，ASML被迫取消20台EUV光刻机订单。中国国产28nm DUV光刻机良率突破85%，计划2025年实现7nm工艺。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        7,
        '2024年美国对华关税调整深度分析',
        '美国政府于2024年5月宣布对价值1800亿美元的中国商品加征关税，涉及电动汽车、锂电池及关键矿产，税率从7.5%提升至25%。此举直接影响特斯拉上海工厂对美出口计划，导致其股价单日下跌4.2%。中国商务部回应将采取反制措施，包括对美农产品加征15%报复性关税。专家预测，全球供应链可能向东南亚转移，越南和墨西哥的制造业投资同比增长23%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        8,
        'GPT-5与多模态AI的突破性进展',
        'OpenAI于2024年3月发布的GPT-5模型，参数规模达到1.8万亿，推理能力逼近人类专家水平。其多模态处理能力支持同时解析文本、图像与音频输入，在医疗诊断测试中，对X光片的异常检测准确率达98.7%。微软已将其集成到Azure云服务，每小时推理成本降低至0.12美元。伦理争议随之而来：欧盟AI法案要求所有生成内容必须添加数字水印。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        9,
        '劳动者必读：10步完成有效劳动仲裁',
        '根据《劳动争议调解仲裁法》，劳动者需在权益受损1年内提交仲裁申请。关键步骤包括：1）收集劳动合同、考勤记录；2）向企业注册地仲裁委提交申请书（模板可在人社部官网下载）；3）准备银行流水证明欠薪事实。2023年数据显示，北京地区仲裁胜诉率68%，平均赔偿金额为3.2个月工资。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        10,
        '拯救发际线：程序员防脱发全攻略',
        '长期熬夜与压力导致程序员脱发率高达43%（《2024程序员健康白皮书》）。建议：1）每日补充5mg生物素+1000IU维生素D；2）使用含咖啡因成分的防脱洗发液（如Alpecin C1）；3）每周进行2次头皮按摩（促进毛囊血液循环）。MIT研究显示，采用间歇式工作法（每编码45分钟休息15分钟）可降低皮质醇水平27%。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        11,
        '高尿酸血症饮食管理：从480μmol/L降到360的实战经验',
        '男性尿酸正常值应＜420μmol/L。实测发现，每日摄入嘌呤需控制在300mg以下：避免动物内脏（猪肝嘌呤含量275mg/100g），选择低脂乳制品。推荐食谱：早餐燕麦粥（嘌呤15mg）+ 午餐清蒸鲈鱼（82mg）+ 晚餐凉拌秋葵（60mg）。配合每日饮水2.5L，可提升尿酸排泄效率40%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        12,
        '2024全球芯片战争：美国对华技术封锁新规',
        '美国商务部5月更新《出口管制条例》，限制14nm以下制程设备对华出口，ASML被迫取消20台EUV光刻机订单。中国国产28nm DUV光刻机良率突破85%，计划2025年实现7nm工艺。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        13,
        '2024年美国对华关税调整深度分析',
        '美国政府于2024年5月宣布对价值1800亿美元的中国商品加征关税，涉及电动汽车、锂电池及关键矿产，税率从7.5%提升至25%。此举直接影响特斯拉上海工厂对美出口计划，导致其股价单日下跌4.2%。中国商务部回应将采取反制措施，包括对美农产品加征15%报复性关税。专家预测，全球供应链可能向东南亚转移，越南和墨西哥的制造业投资同比增长23%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        14,
        'GPT-5与多模态AI的突破性进展',
        'OpenAI于2024年3月发布的GPT-5模型，参数规模达到1.8万亿，推理能力逼近人类专家水平。其多模态处理能力支持同时解析文本、图像与音频输入，在医疗诊断测试中，对X光片的异常检测准确率达98.7%。微软已将其集成到Azure云服务，每小时推理成本降低至0.12美元。伦理争议随之而来：欧盟AI法案要求所有生成内容必须添加数字水印。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        15,
        '劳动者必读：10步完成有效劳动仲裁',
        '根据《劳动争议调解仲裁法》，劳动者需在权益受损1年内提交仲裁申请。关键步骤包括：1）收集劳动合同、考勤记录；2）向企业注册地仲裁委提交申请书（模板可在人社部官网下载）；3）准备银行流水证明欠薪事实。2023年数据显示，北京地区仲裁胜诉率68%，平均赔偿金额为3.2个月工资。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        16,
        '拯救发际线：程序员防脱发全攻略',
        '长期熬夜与压力导致程序员脱发率高达43%（《2024程序员健康白皮书》）。建议：1）每日补充5mg生物素+1000IU维生素D；2）使用含咖啡因成分的防脱洗发液（如Alpecin C1）；3）每周进行2次头皮按摩（促进毛囊血液循环）。MIT研究显示，采用间歇式工作法（每编码45分钟休息15分钟）可降低皮质醇水平27%。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        17,
        '高尿酸血症饮食管理：从480μmol/L降到360的实战经验',
        '男性尿酸正常值应＜420μmol/L。实测发现，每日摄入嘌呤需控制在300mg以下：避免动物内脏（猪肝嘌呤含量275mg/100g），选择低脂乳制品。推荐食谱：早餐燕麦粥（嘌呤15mg）+ 午餐清蒸鲈鱼（82mg）+ 晚餐凉拌秋葵（60mg）。配合每日饮水2.5L，可提升尿酸排泄效率40%。',
        '1',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.articles (
        id,
        title,
        content,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        18,
        '2024全球芯片战争：美国对华技术封锁新规',
        '美国商务部5月更新《出口管制条例》，限制14nm以下制程设备对华出口，ASML被迫取消20台EUV光刻机订单。中国国产28nm DUV光刻机良率突破85%，计划2025年实现7nm工艺。',
        '2',
        '2025-04-09 14:48:16',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        1,
        '国际政策',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        2,
        '人工智能',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        3,
        '法律维权',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        4,
        '程序员健康',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        5,
        '医疗保健',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        6,
        '经济贸易',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        7,
        '职场权益',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        8,
        '技术前沿',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        9,
        '健康饮食',
        '1',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.tags (
        id,
        tag,
        user_id,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        10,
        '软件开发',
        '2',
        '2025-04-09 14:45:45',
        null,
        null
    );

INSERT INTO
    icu.users (
        id,
        name,
        avatar,
        email,
        username,
        signature,
        password,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        1,
        '穿着西服的李白',
        '17389998501824965003600.jpg',
        '',
        'zhubo6984',
        '你到底是谁？',
        '$2a$10$7i3iTmo0vafk9N05/TK1I.6HxUmO5LcOPLlfpWS5wyXRqHh/ja2vO',
        '2025-02-05 14:02:01',
        null,
        null
    );

INSERT INTO
    icu.users (
        id,
        name,
        avatar,
        email,
        username,
        signature,
        password,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        2,
        '北京打手',
        '1744165185002123400627.jpg',
        '',
        'zhubo1997',
        '你到底是谁？',
        '$2a$10$m7KMh6hHPe88yUiDx60scemBuLTD7EAUhPf0YCAbuHAgUW2B2w5cW',
        '2025-02-05 14:42:51',
        null,
        null
    );

INSERT INTO
    icu.users (
        id,
        name,
        avatar,
        email,
        username,
        signature,
        password,
        created_at,
        deleted_at,
        updated_at
    )
VALUES (
        3,
        'ZhuBo_3211',
        '17389998501824965003600.jpg',
        '',
        'zhubo3211',
        '你到底是谁？',
        '$2a$10$hDKVpqlPDD8IBpJPuXEOB.f3atdoGwmW2F..cGZziuaL04eoigLbW',
        '2025-02-08 14:14:47',
        null,
        null
    );