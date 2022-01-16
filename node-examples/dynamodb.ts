import { DynamoDB } from "@aws-sdk/client-dynamodb";

const dynamodb = new DynamoDB({
    region: "ap-northeast-1",
    endpoint: "http://localhost:4566",
});

(async () => {
    // テーブル一覧取得
    const listTablesOutput = await dynamodb.listTables({});
    if (!listTablesOutput.TableNames) {
        throw new Error("listTables failed");
    }
    console.log(listTablesOutput.TableNames);

    const tableName = "todo";

    // Put item
    const putItemOutput = await dynamodb.putItem({
        TableName: tableName,
        Item: {
            id: { S: "0001" },
            title: { S: "記事を書く" },
        },
    });

    // Get item
    const getItemOutput = await dynamodb.getItem({
        TableName: tableName,
        Key: { id: { S: "0001" } },
    });
    if (getItemOutput.Item) {
        console.log(getItemOutput.Item);
    }
})().catch((err) => console.error(err));
