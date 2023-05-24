import NodeBundlr from '@bundlr-network/client/build/cjs/node/bundlr';
import { ContentType, Tag } from '../proto/proto/v1/web3_gateway/storage';
import { ArweaveConfig } from '../config/config';
import { logger } from '../util';

export class ArweaveService {
  private readonly bundler: NodeBundlr;
  private readonly cfg: ArweaveConfig;
  constructor(bundler: NodeBundlr, cfg: ArweaveConfig) {
    this.bundler = bundler;
    this.cfg = cfg;
  }

  public async uploadContent(
    content: Buffer,
    contentType: ContentType,
    tags: Tag[],
  ): Promise<string> {
    try {
      switch (contentType) {
        case ContentType.MARKDOWN: {
          const contentDecoded = content.toString('binary');
          const response = await this.bundler.upload(contentDecoded, { tags });
          return `${this.cfg.base_url}` + `${response.id}`;
        }
        default: {
          logger.error(
            'Unrecognized content type: ',
            `${contentType.toString()}`,
          );
          throw new Error(
            `Unrecognized content type: ${contentType.toString()}`,
          );
        }
      }
    } catch (e) {
      logger.error('failed to upload content to Arweave', e);
      throw new Error(`failed to upload content to Arweave: ${e}`);
    }
  }
}
