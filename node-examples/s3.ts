import { S3 } from "@aws-sdk/client-s3";

const s3 = new S3({
    endpoint: "http://localhost:4566",
});
(async () => {
    // バケット一覧を取得
    const output = await s3.listBuckets({});
    const bucketList = output?.Buckets?.map((bucket) => {
        return bucket.Name;
    });
    console.log(bucketList);
})().catch((err) => {
    console.error(err);
});
