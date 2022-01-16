import { Buffer } from "buffer";

import { KMS } from "@aws-sdk/client-kms";

const endpoint = "http://localhost:4566";
const keyId = "alias/local-kms-key";

const kms = new KMS({
    endpoint,
});

(async () => {
    // Encrypt
    const plainText = "test";
    const encryptOutput = await kms.encrypt({
        KeyId: keyId,
        Plaintext: Buffer.from(plainText),
    });
    if (!encryptOutput.CiphertextBlob) {
        throw new Error("kms.encrypt failed");
    }
    const encryptedBase64 = Buffer.from(encryptOutput.CiphertextBlob).toString("base64");
    console.log(encryptedBase64);

    // Decrypt
    const decryptOutput = await kms.decrypt({
        KeyId: keyId,
        CiphertextBlob: Buffer.from(encryptedBase64, "base64"),
    });
    if (!decryptOutput.Plaintext) {
        throw new Error("kms.decrypt failed");
    }
    const decypted = Buffer.from(decryptOutput.Plaintext).toString("utf8");
    console.log(decypted);
})();
