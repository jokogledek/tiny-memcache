### Overview
<hr/>
Simple library to manage local cache memory

### Notes
```
- for smaller memory usage, use avro encoding instead of just marshaling struct to bytes array
- https://github.com/hamba/avro
- https://avro.apache.org/docs/current/#compare
- use other json encoding library instead of default encoding/json library for faster marshal/unmarshal operation
