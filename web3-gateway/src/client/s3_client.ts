import {
  S3Client,
  GetObjectCommand,
  GetObjectCommandOutput,
} from '@aws-sdk/client-s3';
import { Readable } from 'stream';

export class AWSS3Client {
  private readonly s3Client;

  constructor(region: string) {
    this.s3Client = new S3Client({ region });
  }

  public async downloadS3Content(
    bucket: string,
    key: string,
  ): Promise<{ content: Buffer; contentType: string }> {
    // Create a command to get the object
    const command = new GetObjectCommand({
      Bucket: bucket,
      Key: key,
    });

    try {
      // Send the command to the S3 client and receive the response
      const response: GetObjectCommandOutput = await this.s3Client.send(
        command,
      );

      // Response.Body is a readable stream. Convert it to a Buffer
      const content = await this.streamToBuffer(response.Body as Readable);

      // Extract the content type from the response
      const contentType = response.ContentType ?? 'unknown';

      return { content, contentType };
    } catch (error) {
      // Handle any error that occurred during the API call or stream handling
      console.error('Error downloading content from S3:', error);
      throw error;
    }
  }

  private async streamToBuffer(stream: Readable): Promise<Buffer> {
    return new Promise((resolve, reject) => {
      const chunks: Uint8Array[] = [];
      stream.on('data', (chunk) => chunks.push(chunk));
      stream.once('end', () => resolve(Buffer.concat(chunks)));
      stream.once('error', reject);
    });
  }
}
