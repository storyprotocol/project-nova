## Operation Task Commands:

**Backfill Collection Data**
```
task backfill --collection=<collection address> --start=<start token id> --end=<end token id>
```

**Upload content from Notion to S3**
```
task content-upload --page_id=<notion page id> --s3_path=<s3 bucket>/<s3 folder> --release_at=<2023-03-24T00:00:00Z>
```

**Upload Merkle proof files from S3 to DB**
```
task proof --allowlist=<allowlist id> --json=<json file name in S3>
```

**Open DB Shell**
```
task db-shell
```

**Upgrade DB Schema**
```
task db-up
```