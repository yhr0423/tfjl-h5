'''
没有mongorestore命令时，使用bson文件通过python代码导入数据
'''

from pymongo import MongoClient
import bson
import json
import os

if __name__ == '__main__':
    client = MongoClient('localhost', 27017)
    db = client['tfjl']
    for filename in os.listdir('dump/tfjl'):
        if filename.endswith('.bson'): # 通过bson文件导入数据
            collection = db[os.path.splitext(filename)[0]]
            with open(os.path.join('dump/tfjl', filename), 'rb') as f:
                data = bson.decode_file_iter(f)
                collection.insert_many(data)
        elif filename.endswith('.metadata.json'): # 通过metadata.json文件创建索引
            with open(os.path.join('dump/tfjl', filename)) as f:
                metadata = json.load(f)
                if metadata["type"] == "collection":
                    collection = db[metadata["collectionName"]]
                    indexes = metadata['indexes']
                    for index in indexes:
                        key = index['key']
                        name = index['name']
                        keys = []
                        for item in key:
                            keys.append((item, int(key[item]["$numberInt"])))
                        collection.create_index(keys, name=name)
    client.close()
