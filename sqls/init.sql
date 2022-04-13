/*
 Navicat Premium Data Transfer

 Source Server         : farm
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 13/04/2022 17:03:03
*/

PRAGMA foreign_keys = false;

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
INSERT INTO "conf" VALUES (2, 'WeightMin', 100);
INSERT INTO "conf" VALUES (3, 'HeightMin', 30);
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
INSERT INTO "userinfo" VALUES (1, 'username', '');
INSERT INTO "userinfo" VALUES (2, 'userage', '');
INSERT INTO "userinfo" VALUES (3, 'farmname', '');
INSERT INTO "userinfo" VALUES (4, 'farmclassid', '');
INSERT INTO "userinfo" VALUES (5, 'petname', '');
COMMIT;



PRAGMA foreign_keys = true;
