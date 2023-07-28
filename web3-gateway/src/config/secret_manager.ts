import {
  SecretsManagerClient,
  GetSecretValueCommand,
} from '@aws-sdk/client-secrets-manager';

export class SecretsManager {
  private readonly smClient: SecretsManagerClient;

  constructor(smClient: SecretsManagerClient) {
    this.smClient = smClient;
  }

  public async fetchSecrets(appID: string): Promise<string> {
    try {
      const params = {
        SecretId: appID,
      };
      const data = await this.smClient.send(new GetSecretValueCommand(params));
      return data.SecretString;
    } catch (err) {
      throw new Error(`Failed to fetch secrets ${err}`);
    }
  }
}
