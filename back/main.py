from datetime import datetime

import cv2
import multiprocessing
import time
import os
import uvicorn

import mysql.connector
from fastapi import FastAPI, Query

app = FastAPI()

def record_video_stream(url, queue, code, index, duration=3, num_frames_to_capture=5):
    t1 = time.time()
    cap = cv2.VideoCapture(url)
    print("time", time.time() - t1)
    if not cap.isOpened():
        print(f"Error opening video stream {url}")
        return

    # Create a valid filename from the URL
    folder_name = f"www/wwwroot/likeadmin_go/public/uploads/image/camera_{index}"
    os.makedirs(folder_name, exist_ok=True)
    output_path = os.path.join(os.getcwd(), f"{folder_name}/{code}.mp4")


    # Define the codec and create VideoWriter object
    fps = cap.get(cv2.CAP_PROP_FPS)
    fourcc = cv2.VideoWriter_fourcc(*'mp4v')
    out = cv2.VideoWriter(output_path, fourcc, fps, (int(cap.get(3)), int(cap.get(4))))

    captured_frames = 0
    total_frames = 0
    image_paths = []

    start_time = time.time()
    while (time.time() - start_time) < duration:
        ret, frame = cap.read()
        if not ret:
            break
        out.write(frame)  # Write the frame into the file

        # Capture a frame at the specified interval
        if total_frames % fps == 0 and captured_frames < num_frames_to_capture:
            cv2.imwrite(os.path.join(os.getcwd(), f"{folder_name}/{code}_frame{captured_frames + 1}_{str(int(time.time()))}.jpg"), frame)
            img_path = f"camera_{index}/{code}_frame{captured_frames + 1}_{str(int(time.time()))}.jpg"
            print(img_path)
            image_paths.append("http://127.0.0.1:8000/api/uploads/image/" + img_path)  # 将图片路径添加到列表中
            captured_frames += 1

        total_frames += 1

    # Release everything when the job is finished
    cap.release()
    out.release()
    queue.put(image_paths)

def main(code):
    rtsp_urls = [
        "rtsp://admin:a1234567@192.168.31.67:554/Streaming/Channels/1"
    ]

    frame_queue = multiprocessing.Queue()
    processes = []
    num_cameras = len(rtsp_urls)

    for index, url in enumerate(rtsp_urls):
        process = multiprocessing.Process(target=record_video_stream, args=(url, frame_queue, code, index))
        process.start()
        processes.append(process)

    image_results = [] * num_cameras
    for _ in range(num_cameras):
        image_paths = frame_queue.get()  # This will block until it receives a result
        image_results.append(image_paths)

    # Wait for all processes to complete
    for process in processes:
        process.join()

    return image_results

def insert_image_path_to_database(image_paths_1, image_paths_2, image_paths_3, code):
    # 假设使用 mysql-connector-python 来连接 MySQL 数据库
    connection = mysql.connector.connect(host='127.0.0.1',
                                         port="3306",
                                         database='deviceresource',
                                         user='root',
                                         password='120812614')
    cursor = connection.cursor()
    print(image_paths_1)
    # 准备 SQL 查询语句，将图片路径插入数据库表
    update_query = f"UPDATE la_resource SET img_top = %s, img_front = %s, img_behind = %s WHERE device_code = %s"
    data = (','.join(image_paths_1), ','.join(image_paths_2), ','.join(image_paths_3), code)  # 将列表转换为逗号分隔的字符串

    # 执行 SQL 查询
    cursor.execute(update_query, data)

    # 提交事务
    connection.commit()

    # 关闭数据库连接
    cursor.close()
    connection.close()

@app.get("/")
async def capture_cameras(code: str = Query(...)):
    image_results = main(code)
    if image_results:
        # Assume you have a function to handle the database insertion
        insert_image_path_to_database(image_results[0], image_results[1], [], code)
    return {"status": "success", "data": image_results}

if __name__ == '__main__':
    # 加载预训练的分词器和模型
    uvicorn.run(app, host='0.0.0.0', port=9090, workers=1)
