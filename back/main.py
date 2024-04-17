from fastapi import FastAPI, Response, Query
from starlette.responses import StreamingResponse
import cv2
import mysql.connector
import os
import asyncio
import time

app = FastAPI()

# RTSP地址列表
rtsp_urls = [
    "rtsp://admin:a1234567@192.168.31.68/live",
    "rtsp://admin:a1234567@192.168.31.71:554/Streaming/Channels/1",
]

video_captures = [cv2.VideoCapture(rtsp_url) for rtsp_url in rtsp_urls]

async def generate_frames(cap, camera_index, code):
#     cap = cv2.VideoCapture(rtsp_url)
    frame_count = 0

    # 创建文件夹以保存视频和图片
    folder_name = f"www/wwwroot/likeadmin_go/public/uploads/image/camera_{camera_index}"
    os.makedirs(folder_name, exist_ok=True)
    video_writer = cv2.VideoWriter(f"{folder_name}/{code}.avi", cv2.VideoWriter_fourcc(*'XVID'), cap.get(cv2.CAP_PROP_FPS), (int(cap.get(cv2.CAP_PROP_FRAME_WIDTH)), int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))))

    image_paths = []  # 用于存储图片路径的列表
    T1 = time.time()
    while frame_count < 5 * cap.get(cv2.CAP_PROP_FPS):
        ret, frame = cap.read()
        if not ret:
            break

        # 每秒截取一张图片
        if frame_count % cap.get(cv2.CAP_PROP_FPS) == 0:
            image_path = f"{folder_name}/{code}_{frame_count // cap.get(cv2.CAP_PROP_FPS)}.jpg"
            cv2.imwrite(image_path, frame, [cv2.IMWRITE_JPEG_QUALITY, 50])
            img_path = f"camera_{camera_index}/{code}_{frame_count // cap.get(cv2.CAP_PROP_FPS)}.jpg"
            print(img_path)
            image_paths.append("http://127.0.0.1:8000/api/uploads/image/"+img_path)  # 将图片路径添加到列表中

        # 保存视频
        video_writer.write(frame)
        frame_count += 1
        await asyncio.sleep(0.0005 / cap.get(cv2.CAP_PROP_FPS))
    T2 = time.time()
    print((T2-T1))
    return image_paths  # 返回图片路径列表


def insert_image_path_to_database(image_paths_1, image_paths_2, code):
    # 假设使用 mysql-connector-python 来连接 MySQL 数据库
    connection = mysql.connector.connect(host='127.0.0.1',
                                         port="3306",
                                         database='deviceresource',
                                         user='root',
                                         password='120812614')
    cursor = connection.cursor()
    print(image_paths_1)
    # 准备 SQL 查询语句，将图片路径插入数据库表
    update_query = f"UPDATE la_resource SET img_top = %s, img_front = %s WHERE device_code = %s"
    data = (','.join(image_paths_1), ','.join(image_paths_2), code)  # 将列表转换为逗号分隔的字符串

    # 执行 SQL 查询
    cursor.execute(update_query, data)

    # 提交事务
    connection.commit()

    # 关闭数据库连接
    cursor.close()
    connection.close()


@app.get("/")
async def get_video_feed(response: Response, code: str = Query(...)):
    # 设置响应头，指示返回的是multipart格式的数据
    response.headers['Content-Type'] = 'multipart/x-mixed-replace; boundary=frame'

    # 捕获每个摄像头的视频流
    async def stream_multiple_cameras():
        tasks = [asyncio.create_task(generate_frames(cap, i, code)) for i, cap in enumerate(video_captures)]
        # 并行执行多个异步任务
        image_paths_lists = await asyncio.gather(*tasks)
        # 将图片路径列表传递给插入数据库函数
        insert_image_path_to_database(image_paths_lists[0], image_paths_lists[1], code)

    # 执行异步任务
    await stream_multiple_cameras()

    # 返回空白数据流，因为数据已经被写入数据库，不需要再返回图片数据
    return StreamingResponse(iter([]))