# IEOR142-EV-registration-prediction
# Goal
The goal of this project is to predict the total registration number of electric vehicles in washington. 
# Method
We intend to use **linear regression , Cart, Random Forest and ARIMA models**.
# Steps
## Step 1 Find the data of EV registration (dependent variable)
data processed: The total registration of EV in washington each week for each zipcode
[https://docs.google.com/spreadsheets/d/1yvA6QSGX82xFxjYSa_GkTy_do4KOC22-9XSZPq0CU8E/edit#gid=885632676]


## Step 2 Find the data of independent variables
Data should be put inside [https://docs.google.com/spreadsheets/d/1yvA6QSGX82xFxjYSa_GkTy_do4KOC22-9XSZPq0CU8E/edit#gid=0] (raw data description)

**Possible independent variable**

Household-related attributes: household size, income, location


Trip attributes: daily trips


Enegry: Price differential between gasoline and electricity, CPIenergy


Average gasoline price in the U.S.: https://tradingeconomics.com/commodity/gasoline


Average Lithium-ion battery pack prices, $/kWh: https://tradingeconomics.com/commodity/lithium 


State/federal incentive events


Number of public-registered EV charging stations: https://openchargemap.org/site/develop/api




## Step3 Clean data and generate excel/csv file
See [https://docs.google.com/spreadsheets/d/1yvA6QSGX82xFxjYSa_GkTy_do4KOC22-9XSZPq0CU8E/edit#gid=0] ( sheet name processed data)
each independent variable should be added as a new column

## Setp4 Train data with different models



