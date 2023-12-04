import {
  sendUnaryData,
  ServerUnaryCall,
  status,
  UntypedHandleCall,
} from '@grpc/grpc-js';
import {
  StorageServiceServer,
  StorageServiceService,
} from '../proto/proto/v1/web3_gateway/service';
import { ArweaveService } from '../service/arweave_service';
import {
  FundIrysReq,
  FundIrysResp,
  StorageType,
  UploadContentReq,
  UploadContentResp,
} from '../proto/proto/v1/web3_gateway/storage';
import { ServiceError, logger } from '../util';
import { AWSS3Client } from 'src/client/s3_client';

export class StorageServer implements StorageServiceServer {
  [method: string]: UntypedHandleCall;

  private readonly arweaveService;
  private readonly s3Client;

  constructor(arweaveService: ArweaveService, s3Client: AWSS3Client) {
    this.arweaveService = arweaveService;
    this.s3Client = s3Client;
  }

  public async fundIrysAccount(
    call: ServerUnaryCall<FundIrysReq, FundIrysResp>,
    callback: sendUnaryData<FundIrysResp>,
  ): Promise<any> {
    const { amountInWei } = call.request;
    const fundIrysRes = await (
      this.arweaveService as ArweaveService
    ).fundIrysAccount(amountInWei);
    callback(null, FundIrysResp.fromJSON(fundIrysRes));
  }

  public async uploadContent(
    call: ServerUnaryCall<UploadContentReq, UploadContentResp>,
    callback: sendUnaryData<UploadContentResp>,
  ): Promise<any> {
    try {
      const { storage, content, contentType, tags, s3Bucket, s3Key } =
        call.request;
      switch (storage) {
        case StorageType.ARWEAVE: {
          let contentUrl;
          if (s3Bucket && s3Key) {
            const { content, contentType } =
              await this.s3Client.downloadS3Content(s3Bucket, s3Key);
            contentUrl = await (
              this.arweaveService as ArweaveService
            ).uploadContent(content, contentType, tags);
          } else if (content && contentType) {
            contentUrl = await (
              this.arweaveService as ArweaveService
            ).uploadContent(content, contentType, tags);
          }
          logger.info('Uploaded content to Arweave: ', `${contentUrl}`, tags);
          callback(
            null,
            UploadContentResp.fromJSON({
              contentUrl: contentUrl,
            }),
          );
          break;
        }
        default: {
          logger.error('Unrecognized storage type: ', `${storage.toString()}`);
          throw new Error(`Unrecognized storage type: ${storage.toString()}`);
        }
      }
    } catch (err) {
      callback(
        new ServiceError(status.INTERNAL, `Failed to upload content : ${err}`),
        null,
      );
    }
  }
}

export const StorageGrpcService = StorageServiceService;
