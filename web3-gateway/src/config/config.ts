import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { join } from 'path';
import { plainToInstance, plainToClassFromExist } from 'class-transformer';
import {
  ValidateNested,
  validateOrReject,
  IsIn,
  Equals,
  IsNotEmpty,
} from 'class-validator';
import { SecretsManager } from './secret_manager';

const YAML_CONFIG_FILENAME = './base.yaml';

export enum Env {
  Prod = 'prod',
  Staging = 'staging',
  Dev = 'dev',
  Local = 'local',
}

class ServerConfig {
  @IsNotEmpty()
  port: string;
}

export class ArweaveConfig {
  @IsNotEmpty()
  base_url: string;

  @IsNotEmpty()
  bundlr_node: string;

  @IsNotEmpty()
  wallet_chain: string;
}

export class Config {
  @IsIn(['local', 'dev', 'staging', 'prod'])
  env: string;

  @Equals('web3-gateway')
  app_id: string;

  @ValidateNested()
  server: ServerConfig;

  @IsNotEmpty()
  wallet_key: string;

  @ValidateNested()
  arweave_config: ArweaveConfig;
}

export const initializeConfig = async (): Promise<Config> => {
  try {
    const baseCfg = yaml.load(
      readFileSync(join(__dirname, YAML_CONFIG_FILENAME), 'utf8'),
    ) as Record<string, any>;

    const env = process.env.SERVER_ENV;
    const extendCfg = yaml.load(
      readFileSync(join(__dirname, `${env}.yaml`), 'utf8'),
    ) as Record<string, any>;

    let mergedCfg = plainToInstance(Config, { ...baseCfg, ...extendCfg });
    if (mergedCfg.env !== Env.Local) {
      const secret = await SecretsManager.fetchSecrets(mergedCfg.app_id);
      mergedCfg = plainToClassFromExist(mergedCfg, JSON.parse(secret));
    }

    validateOrReject(mergedCfg).catch((errors) => {
      throw new Error(`Config validation failed. Errors: ${errors}`);
    });

    return Promise.resolve(mergedCfg);
  } catch (err) {
    throw new Error(`Failed to init config ${err}`);
  }
};
