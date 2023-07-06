
import os
import base64
import json
import asyncio
import aiohttp

protocolNumMap = {
    1000001: "登录验证",
    1000002: "登录请求角色列表",
    1000003: "登录请求创建角色",
    1000004: "登录请求选择角色",
    1000005: "登录验证在线",
    1000006: "登录验证Ping",
    2000001: "角色-角色进入游戏",
    2000002: "角色-同步角色数据",
    2000003: "角色-同步角色属性",
    2000004: "角色-同步角色信息",
    2000007: "角色-同步角色邮件数据",
    2000009: "角色-同步角色任务数据",
    2000011: "角色-同步战斗数据",
    2000025: "角色-设置默认战斗阵容",
    2000026: "角色-设置战斗阵容",
    2000030: "角色-角色对战结算数据",
    2000034: "角色-设置引导步骤",
    2000039: "角色-角色简要信息",
    2000052: "角色-同步外挂数据",
    2000054: "角色-角色总看广告数据",
    2000071: "角色-同步角色任务额外数据",
    2000105: "角色-同步角色条件数据",
    2000121: "角色-同步角色开关数据",
    2000131: "角色-同步角色消耗数据",
    2000138: "角色-修改角色战车皮肤",
    2000151: "角色-修改角色英雄皮肤",
    2000155: "角色-同步角色章节数据",
    4000001: "活动-同步角色所有活动数据",
    4000014: "活动-获取颁奖日数据",
    4000016: "活动-获取黄金联赛数据",
    4000047: "活动-同步角色试炼场数据",
    4000050: "活动-同步角色公会战数据",
    4000052: "活动-获取角色大航海数据",
    4000053: "活动-同步角色大航海数据",
    4000057: "活动-同步角色天空之城数据",
    4000060: "活动-同步角色寒冰堡数据",
    4000063: "活动-同步角色周年庆数据",
    4000065: "活动-同步角色回归数据",
    4000067: "活动-同步角色雾隐数据",
    4000069: "活动-机械迷城数据",
    6000001: "聊天请求",
    6000003: "战斗匹配房间关闭",
    7000001: "同步朋友数据",
    9000001: "战斗匹配",
    9000002: "战斗匹配取消",
    9000003: "战斗匹配结果",
    9000004: "战斗匹配竞争战斗",
    9000005: "战斗匹配对战取消",
	10000002: "对战登录",
	10000003: "对战加载准备",
	10000004: "对战开始",
	10000005: "对战结束",
	10000007: "怪物血量同步",
	10000014: "战斗对战结束提交（多人，结束）",
    10000015: "战斗匹配对战结束(单人，结束)",
    10000020: "提交对战每阶段逻辑数据（多人，过程）",
    10000021: "提交对战每阶段逻辑数据（单人，过程）",
    10000101: "更新英雄",
    10000104: "战斗银币同步",
	10000105: "战斗卖出英雄同步",
    10000107: "战斗英雄属性同步",
    10000108: "卡牌刷新次数同步",
	10000109: "操作装备同步",
    12000001: "同步联盟数据",
	12000033: "获取机械迷城数据",
    12000034: "同步机械迷城数据",
	12000037: "机械迷城选择卡组",
}

headers = {"Authorization": "e756795a-1245-458f-ae1c-8f1e2ccf5e28"}

async def process_file(session, file_path):
    with open(file_path, "rb") as f:
        # Read the content of the file into a bytes object
        content = f.read()

        # Extract the specific range (5th byte to 8th byte) from the content
        extracted_bytes = content[4:8]

        # Decode the extracted bytes using little-endian byte order
        decoded_value = int.from_bytes(extracted_bytes, byteorder="little")

        print(f"Extracted value from {file_path}: {decoded_value}")

        base64Str = base64.b64encode(content[8:]).decode("utf-8")
        # print("base64Str:", base64Str)

        data = {}
        if "client" in file_path:
            # 创建由JSON数据组成的请求正文
            data = {
                "clienttype": 1,
                "protocolnum": decoded_value,
                "bytes": base64Str,
            }
        elif "server" in file_path:
            # 创建由JSON数据组成的请求正文
            data = {
                "clienttype": 2,
                "protocolnum": decoded_value,
                "bytes": base64Str,
            }

        json_data = json.dumps(data)

        async with session.post("http://localhost:8080/tfjlh5/decode", data=json_data, headers=headers) as resp:
            # print(await resp.text())
            return os.path.basename(file_path), decoded_value, await resp.json()

async def read_files(directory):
    async with aiohttp.ClientSession() as session:
        try:
            file_path = os.path.join(directory, "decode.txt")
            os.remove(file_path)
            print(f"File '{file_path}' has been deleted.")
        except OSError as error:
            print(error)

        # Traverse through the provided directory
        tasks = []
        limits = (0, 50000)
        count = 0
        for root, _, files in os.walk(directory):
            for file_name in files:
                if file_name.endswith(".bin"):
                    if file_name.endswith("_000002_server.bin"):
                        await process_file(session, os.path.join(root, file_name))
                    elif count >= limits[0] and count <= limits[1]:
                        task = asyncio.create_task(process_file(session, os.path.join(root, file_name)))
                        tasks.append(task)
                        pass
                    count += 1

        results = await asyncio.gather(*tasks)
        # print(results)
        with open(os.path.join(directory, "decode.txt"), "a", encoding="utf-8") as f:
            for result in results:
                if result[1] in protocolNumMap:
                    f.write(f"{result[0]} {protocolNumMap[result[1]]}: {result[2]}\n")
                else:
                    f.write(f"{result[0]} {result[1]}: {result[2]}\n")

asyncio.run(read_files(r"websocket抓包目录"))
