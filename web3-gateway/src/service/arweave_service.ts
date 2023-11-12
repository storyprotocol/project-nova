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

  /**
   * Uploads content to Arweave.
   * @param {Buffer} content - The content to upload.
   * @param {string} contentType - The MIME type of the content.
   * @param {Tag[]} tags - The tags to add to the content.
   * @returns {Promise<string>} The URL of the uploaded content.
   * @throws {Error} If the upload fails.
   */
  public async uploadContent(
    content: Buffer,
    contentType: string,
    tags: Tag[] = [],
  ): Promise<string> {
    if (!content) {
      throw new Error('Content is required');
    }

    if (!contentType) {
      throw new Error('Content type is required');
    }

    // don't add Content-type tag if it already exists
    const hasContentTypeTag = tags.some(
      (t) => t.name.toLowerCase() === 'content-type',
    );

    tags = hasContentTypeTag
      ? tags
      : [{ name: 'Content-Type', value: contentType }, ...tags];

    try {
      const response = await this.bundler.upload(content, { tags });
      return `${this.cfg.base_url}${response.id}`;
    } catch (e) {
      logger.error('Failed to upload content to Arweave', e);
      throw new Error(`Failed to upload content to Arweave: ${e.message}`);
    }
  }
}
