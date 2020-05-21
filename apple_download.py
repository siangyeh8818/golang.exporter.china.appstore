import time
import os
import sys
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
import csv

url_apple_download = 'http://www.ggkjplus.com/admin/#/login'
chromedriver = '/usr/local/share/chromedriver'

account_apple = os.environ.get('APPLE_ACCOUNT')
password_apple = os.environ.get('APPLE_PASSWORD')

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--no-sandbox')
options.add_argument('--disable-gpu')
options.add_argument('blink-settings=imagesEnabled=false')
options.add_argument('--enable-features=OverlayScrollbar')

driver = webdriver.Chrome(executable_path=chromedriver, options=options)
driver.implicitly_wait(60)

driver.get(url_apple_download);
driver.find_element_by_xpath('/html/body/div/div[2]/div[2]/form/div[2]/div/div/input').send_keys(account_apple)
driver.find_element_by_xpath('/html/body/div/div[2]/div[2]/form/div[3]/div/div/input').send_keys(password_apple)
driver.find_element_by_xpath('/html/body/div/div[2]/div[2]/form/button/span').click()

time.sleep(3)
driver.get('https://www.ggkjplus.com/admin/#/webpackage')
time.sleep(3)
result = driver.find_element_by_xpath('/html/body/div/div[3]/div[2]/section/div/div[1]/ul/li[1]/span/div/span').text
#print('balance: ',result)
with open('output.csv', 'w', newline='') as csvfile:
  writer = csv.writer(csvfile)
  writer.writerow([result])

driver.quit()
