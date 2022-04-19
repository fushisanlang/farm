/*
 Navicat Premium Data Transfer

 Source Server         : farm
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 19/04/2022 17:28:29
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for bag
-- ----------------------------
DROP TABLE IF EXISTS "bag";
CREATE TABLE "bag" (
  "keyid" INTEGER NOT NULL DEFAULT 0,
  "id" INTEGER NOT NULL,
  "bagid" INTEGER NOT NULL DEFAULT 0,
  "linkid" integer DEFAULT 0,
  "groupid" INTEGER,
  "countnum" INTEGER
);

-- ----------------------------
-- Records of bag
-- ----------------------------
BEGIN;
INSERT INTO "bag" VALUES (1, 1, 1, 3, 1, 0);
INSERT INTO "bag" VALUES (2, 1, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 1, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 1, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 1, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 1, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 1, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 1, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 1, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 1, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (11, 2, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (12, 2, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (13, 2, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (14, 2, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (15, 2, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (16, 2, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (17, 2, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (18, 2, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (19, 2, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (20, 2, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (21, 3, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (22, 3, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (23, 3, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (24, 3, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (25, 3, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (26, 3, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (27, 3, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (28, 3, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (29, 3, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (30, 3, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (31, 4, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (32, 4, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (33, 4, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (34, 4, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (35, 4, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (36, 4, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (37, 4, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (38, 4, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (39, 4, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (40, 4, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (41, 5, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (42, 5, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (43, 5, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (44, 5, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (45, 5, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (46, 5, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (47, 5, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (48, 5, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (49, 5, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (50, 5, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (51, 6, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (52, 6, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (53, 6, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (54, 6, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (55, 6, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (56, 6, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (57, 6, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (58, 6, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (59, 6, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (60, 6, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (61, 7, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (62, 7, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (63, 7, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (64, 7, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (65, 7, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (66, 7, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (67, 7, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (68, 7, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (69, 7, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (70, 7, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (71, 8, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (72, 8, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (73, 8, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (74, 8, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (75, 8, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (76, 8, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (77, 8, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (78, 8, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (79, 8, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (80, 8, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (81, 9, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (82, 9, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (83, 9, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (84, 9, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (85, 9, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (86, 9, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (87, 9, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (88, 9, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (89, 9, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (90, 9, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (91, 10, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (92, 10, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (93, 10, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (94, 10, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (95, 10, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (96, 10, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (97, 10, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (98, 10, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (99, 10, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (100, 10, 10, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for conf
-- ----------------------------
DROP TABLE IF EXISTS "conf";
CREATE TABLE "conf" (
                        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                        "confKey" TEXT NOT NULL,
                        "confValue" TEXT
);

-- ----------------------------
-- Records of conf
-- ----------------------------
BEGIN;
INSERT INTO "conf" VALUES (1, 'version', '0.0.1');
INSERT INTO "conf" VALUES (2, 'WeightMin', 120);
INSERT INTO "conf" VALUES (3, 'HeightMin', 40);
COMMIT;

-- ----------------------------
-- Table structure for farmclass
-- ----------------------------
DROP TABLE IF EXISTS "farmclass";
CREATE TABLE "farmclass" (
                             "farmclassid" INTEGER NOT NULL,
                             "farmclassname" TEXT NOT NULL,
                             "farmclassinfo" TEXT NOT NULL,
                             PRIMARY KEY ("farmclassid")
);

-- ----------------------------
-- Records of farmclass
-- ----------------------------
BEGIN;
INSERT INTO "farmclass" VALUES (1, '昆仑山', '雨水多');
INSERT INTO "farmclass" VALUES (2, '长白山', '气候寒冷');
INSERT INTO "farmclass" VALUES (3, '武当山', '均衡');
INSERT INTO "farmclass" VALUES (4, '南疆', '虫子多');
COMMIT;

-- ----------------------------
-- Table structure for field
-- ----------------------------
DROP TABLE IF EXISTS "field";
CREATE TABLE "field" (
  "keyid" INTEGER NOT NULL DEFAULT 0,
  "id" INTEGER NOT NULL,
  "fieldid" integer NOT NULL,
  "plantid" INTEGER NOT NULL,
  "duringtime" float NOT NULL,
  "isopen" integer NOT NULL
);

-- ----------------------------
-- Records of field
-- ----------------------------
BEGIN;
INSERT INTO "field" VALUES (1, 1, 1, 0, 0.0, 1);
INSERT INTO "field" VALUES (2, 1, 2, 3, 546.0, 1);
INSERT INTO "field" VALUES (3, 1, 3, 3, 546.0, 1);
INSERT INTO "field" VALUES (4, 1, 4, 0, 0.0, 1);
INSERT INTO "field" VALUES (5, 1, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 1, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 1, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 1, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 1, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 1, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (11, 1, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (12, 1, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (13, 2, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (14, 2, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (15, 2, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (16, 2, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (17, 2, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (18, 2, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (19, 2, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (20, 2, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (21, 2, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (22, 2, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (23, 2, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (24, 2, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (25, 3, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (26, 3, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (27, 3, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (28, 3, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (29, 3, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (30, 3, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (31, 3, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (32, 3, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (33, 3, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (34, 3, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (35, 3, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (36, 3, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (38, 4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (49, 5, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (50, 5, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (51, 5, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (52, 5, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (53, 5, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (54, 5, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (55, 5, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (56, 5, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (57, 5, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (58, 5, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (59, 5, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (60, 5, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (61, 6, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (62, 6, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (63, 6, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (64, 6, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (65, 6, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (66, 6, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (67, 6, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (68, 6, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (69, 6, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (70, 6, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (71, 6, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (72, 6, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (73, 7, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (74, 7, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (75, 7, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (76, 7, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (77, 7, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (78, 7, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (79, 7, 7, 0, 0.0, 1);
INSERT INTO "field" VALUES (80, 7, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (81, 7, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (82, 7, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (83, 7, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (84, 7, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (85, 8, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (86, 8, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (87, 8, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (88, 8, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (89, 8, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (90, 8, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (91, 8, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (92, 8, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (93, 8, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (94, 8, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (95, 8, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (96, 8, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (97, 9, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (98, 9, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (99, 9, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (100, 9, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (101, 9, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (102, 9, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (103, 9, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (104, 9, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (105, 9, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (106, 9, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (107, 9, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (108, 9, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (109, 10, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (110, 10, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (111, 10, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (112, 10, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (113, 10, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (114, 10, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (115, 10, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (116, 10, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (117, 10, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (118, 10, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (119, 10, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (120, 10, 12, 0, 0.0, 0);
COMMIT;

-- ----------------------------
-- Table structure for plant
-- ----------------------------
DROP TABLE IF EXISTS "plant";
CREATE TABLE "plant" (
  "id" INTEGER NOT NULL,
  "plantname" TEXT NOT NULL,
  "timeneed" float,
  "level" integer,
  "price" integer,
  "ex" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of plant
-- ----------------------------
BEGIN;
INSERT INTO "plant" VALUES (1, '天仙子', 10000.0, 0, 100, 3);
INSERT INTO "plant" VALUES (2, '幽香绮罗仙品', 200000.0, 20, 100000, 1000);
INSERT INTO "plant" VALUES (3, '天仙子2', 1000.0, 0, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for seed
-- ----------------------------
DROP TABLE IF EXISTS "seed";
CREATE TABLE "seed" (
  "id" INTEGER NOT NULL,
  "seedname" TEXT NOT NULL,
  "plantid" integer,
  "level" integer,
  "price" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of seed
-- ----------------------------
BEGIN;
INSERT INTO "seed" VALUES (1, '天仙子种子', 1, 1, 1000);
INSERT INTO "seed" VALUES (2, '幽香绮罗仙品', 2, 10, 1000);
COMMIT;

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);

-- ----------------------------
-- Records of sqlite_sequence
-- ----------------------------
BEGIN;
INSERT INTO "sqlite_sequence" VALUES ('conf', 3);
INSERT INTO "sqlite_sequence" VALUES ('userinfo', 9);
COMMIT;

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS "userinfo";
CREATE TABLE "userinfo" (
                            "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                            "infoKey" TEXT NOT NULL,
                            "infoValue" TEXT
);

-- ----------------------------
-- Records of userinfo
-- ----------------------------
BEGIN;
INSERT INTO "userinfo" VALUES (1, 'username', 1);
INSERT INTO "userinfo" VALUES (2, 'userage', 2);
INSERT INTO "userinfo" VALUES (3, 'farmname', 3);
INSERT INTO "userinfo" VALUES (4, 'farmclassid', 3);
INSERT INTO "userinfo" VALUES (5, 'petname', 5);
INSERT INTO "userinfo" VALUES (6, 'money', 4000);
INSERT INTO "userinfo" VALUES (7, 'ex', 0);
INSERT INTO "userinfo" VALUES (8, 'level', 0);
INSERT INTO "userinfo" VALUES (9, 'nextfieldneedmoney', 2000);
COMMIT;

-- ----------------------------
-- Auto increment value for conf
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 3 WHERE name = 'conf';

-- ----------------------------
-- Auto increment value for userinfo
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 9 WHERE name = 'userinfo';

PRAGMA foreign_keys = true;
