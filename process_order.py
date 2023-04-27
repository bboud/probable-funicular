import requests
import os
from concurrent.futures import ThreadPoolExecutor
import pyart

ORDER_NUMBER    = 'HAS012373877'
RADAR_CODE      = 'KTLX'
PRODUCT_TYPES   = ['N0Q', 'N0U']

ORDER_URL = f'https://www.ncei.noaa.gov/pub/has/{ORDER_NUMBER}/'
req = requests.get(f'{ORDER_URL}/fileList.txt')
cont = req.content.splitlines()

def download(p_type, c):
    if p_type.encode() in c:
        c_str = c.decode()
        print(f'Downloading {c_str}')
        c_str = c_str.split('/')[1]

        try:
            os.mkdir(f'{RADAR_CODE}_{p_type}')
        except FileExistsError:
            pass
        
        try:
            f = open(f'{RADAR_CODE}_{p_type}/{c_str}', 'x')
            f.close()
        except FileExistsError:
            pass

        try:
            f = open(f'{RADAR_CODE}_{p_type}/{c_str}', 'wb')
            f.write(requests.get(f'{ORDER_URL}{c.decode()}').content)
            f.close()
        except Exception as e:
            print(str(e))

        try:
            radar = pyart.io.read(f'{RADAR_CODE}_{p_type}/{c_str}')
            data = radar.fields['velocity']['data']
        except Exception as e:
            print(str(e))
        
with ThreadPoolExecutor(max_workers=20) as pool:
    for p_type in PRODUCT_TYPES:
        for c in cont:
            pool.submit(download, p_type, c)