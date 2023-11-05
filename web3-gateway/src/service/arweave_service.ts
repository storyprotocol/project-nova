import NodeBundlr from '@bundlr-network/client/build/cjs/node/bundlr';
import { Tag } from '../proto/proto/v1/web3_gateway/storage';
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
    contentType: string,
    tags: Tag[],
  ): Promise<string> {
    // don't add Content-type tag if it already exists
    const hasContentTypeTag =
      tags && tags.some((t) => t.name.toLowerCase() === 'content-type');

    tags = hasContentTypeTag
      ? tags
      : [{ name: 'Content-Type', value: contentType }, ...(tags ?? [])];

    try {
      const response = await this.bundler.upload(content, { tags });
      return `${this.cfg.base_url}` + `${response.id}`;
    } catch (e) {
      logger.error('failed to upload content to Arweave', e);
      throw new Error(`failed to upload content to Arweave: ${e}`);
    }
  }
}
