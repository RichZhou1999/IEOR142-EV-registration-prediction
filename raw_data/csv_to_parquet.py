import pandas as pd
import pyarrow as pa
import pyarrow.parquet as pq

csv_file = './wa_registration_data_processed.csv' 
parquet_file = './wa_registration_data_processed.csv.parquet'
chunksize = 100_000

csv_stream = pd.read_csv(csv_file, sep='\t', chunksize=chunksize, low_memory=False)

for i, chunk in enumerate(csv_stream):
    print("Chunk", i)
    if i == 0:
        parquet_schema = pa.Table.from_pandas(df=chunk).schema
        parquet_writer = pq.ParquetWriter(parquet_file, parquet_schema, compression='snappy')
    table = pa.Table.from_pandas(chunk, schema=parquet_schema)
    parquet_writer.write_table(table)

parquet_writer.close()