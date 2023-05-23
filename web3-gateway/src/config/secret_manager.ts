import {
  SecretsManagerClient,
  GetSecretValueCommand,
} from '@aws-sdk/client-secrets-manager';

const REGION = 'us-east-1';

export class SecretsManager {
  private static readonly smClient: SecretsManagerClient =
    new SecretsManagerClient({ region: REGION });

  public static async fetchSecrets(appID: string): Promise<string> {
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
