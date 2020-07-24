#!/usr/bin/env python
# coding: utf-8

# In[5]:


import time
import os
import sys
import csv
from bs4 import BeautifulSoup

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By


# In[6]:


CHROMEDRIVER_PATH = '/usr/local/share/chromedriver'

account_apple = os.environ.get('APPLE_ACCOUNT')
password_apple = os.environ.get('APPLE_PASSWORD')

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--no-sandbox')
options.add_argument('--disable-gpu')
options.add_argument('blink-settings=imagesEnabled=false')
options.add_argument('--enable-features=OverlayScrollbar')

driver = webdriver.Chrome(executable_path=CHROMEDRIVER_PATH, options=options)
driver.implicitly_wait(60)

driver.get("https://www.ggkjplus.com/admin/#/login")
time.sleep(2)
driver.find_element(By.NAME, "mobile").click()
driver.find_element(By.NAME, "mobile").send_keys(account_apple)
driver.find_element(By.NAME, "pwd").click()
driver.find_element(By.NAME, "pwd").send_keys(password_apple)
driver.find_element(By.CSS_SELECTOR, ".over-btn > span").click()
time.sleep(5)
soup = BeautifulSoup(driver.page_source,"html.parser")
driver.close()

result = int(soup.find_all("div",class_="iCountUp num")[3].text.replace(",",""))
#print("balance: ",result)

with open('output.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow([result])


# In[ ]:




