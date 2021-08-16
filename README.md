# is-it-scotland

1. Get the tool
   ```
   $ go get github.com/hawx/is-it-scotland
   ```

1. Download [OS Code-Point® Open](https://osdatahub.os.uk/downloads/open/CodePointOpen) and unzip somewhere

1. Build a nice small CSV to query from the data
   ```
   $ is-it-scotland --build < PATH/TO/CSVS/*.csv > data.csv
   ```

1. Query the data
   ```
   $ is-it-scotland --dataset data.csv
   Enter postcode to query, CTRL+C to quit:
    > sw1A 2aa
   => no

    > eh99 1sp
   => yes
   ```
