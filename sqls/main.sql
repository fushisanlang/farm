/*
 Navicat Premium Data Transfer

 Source Server         : farm
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 18/04/2022 17:53:17
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for bag
-- ----------------------------
DROP TABLE IF EXISTS "bag";
CREATE TABLE "bag" (
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
INSERT INTO "bag" VALUES (1, 1, 1, 1, 3);
INSERT INTO "bag" VALUES (1, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (1, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (2, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (3, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (4, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (5, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (6, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (7, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (8, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (9, 10, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 1, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 2, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 3, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 4, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 5, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 6, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 7, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 8, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 9, 0, 0, 0);
INSERT INTO "bag" VALUES (10, 10, 0, 0, 0);
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
INSERT INTO "field" VALUES (1, 1, 0, 0.0, 1);
INSERT INTO "field" VALUES (1, 2, 0, 0.0, 1);
INSERT INTO "field" VALUES (1, 3, 0, 0.0, 1);
INSERT INTO "field" VALUES (1, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (1, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (2, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (3, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (4, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (5, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (6, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (7, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (8, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 7, 0, 0.0, 1);
INSERT INTO "field" VALUES (9, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (9, 12, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 1, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 2, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 3, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 4, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 5, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 6, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 7, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 8, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 9, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 10, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 11, 0, 0.0, 0);
INSERT INTO "field" VALUES (10, 12, 0, 0.0, 0);
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
INSERT INTO "plant" VALUES (3, '天仙子2', 1.0, 0, 1, 1);
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
INSERT INTO "userinfo" VALUES (6, 'money', 500);
INSERT INTO "userinfo" VALUES (7, 'ex', 0);
INSERT INTO "userinfo" VALUES (8, 'level', 0);
INSERT INTO "userinfo" VALUES (9, 'nextfieldneedmoney', 1000);
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
