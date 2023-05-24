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
  StorageType,
  UploadContentReq,
  UploadContentResp,
} from '../proto/proto/v1/web3_gateway/storage';
import { ServiceError, logger } from '../util';

export class StorageServer implements StorageServiceServer {
  [method: string]: UntypedHandleCall;

  private readonly arweaveService;

  constructor(arweaveService: ArweaveService) {
    this.arweaveService = arweaveService;
  }

  public async uploadContent(
    call: ServerUnaryCall<UploadContentReq, UploadContentResp>,
    callback: sendUnaryData<UploadContentResp>,
  ): Promise<any> {
    try {
      const { storage, content, contentType, tags } = call.request;
      switch (storage) {
        case StorageType.ARWEAVE: {
          const contentUrl = await (
            this.arweaveService as ArweaveService
          ).uploadContent(content, contentType, tags);
          console.log(contentUrl);
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
