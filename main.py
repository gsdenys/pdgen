import os
import sys

import docx

import pandas as pd
from pgsql import data
from word import to_word as export_to_word


class DataDict:
    def __init__(self) -> None:
        self.__data_dict = {}
        self.__prepare()
        
    def __prepare(self):
        df = data()

        tables = df['Table'].unique()
        
        for table in tables:
            self.__data_dict[table] = df[df['Table'] == table]
            self.__data_dict[table] = self.__data_dict[table].set_index('Colonne')
            del self.__data_dict[table]['Table']
    
    def to_excel(self, file_name: str):
        with pd.ExcelWriter(file_name) as writer:
            for k, v in self.__data_dict.items():
                v.to_excel(writer, sheet_name=k)
                
    def to_world(self):
        export_to_word(self.__data_dict)


def __read_connection():
    file1 = open('connection.db', 'r')
    lines = file1.readline()
    file1.close()
    
    os.environ['DATABASE_URL'] = lines
    
                
if __name__ == "__main__":
    
    __file_name = 'data-dictionary'
    __command = sys.argv[1]
    
    if __command == "connect":
        with open('connection.db', 'w') as file:
            file.write(sys.argv[2])
            
    elif sys.argv[1] == "excel":
        __read_connection()
        
        dt = DataDict()
        __file_name += '.xlsx'
        dt.to_excel(__file_name)
    elif sys.argv[1] == "world":
        __read_connection()
        
        dt = DataDict()
        dt.to_world()
    