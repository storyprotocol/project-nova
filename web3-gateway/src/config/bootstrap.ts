import { Server, ServerCredentials } from '@grpc/grpc-js';
import { addReflection } from 'grpc-server-reflection';
import { StorageServer, StorageGrpcService } from '../server/storage_server';
import { ArweaveService } from '../service/arweave_service';
import NodeBundlr from '@bundlr-network/client';
import { initializeConfig } from './config';
import { logger } from '../util';

export const bootstrap = async () => {
  const proto_descriptor_path = 'dist/proto/proto_descriptor.bin';
  const server = new Server({
    'grpc.max_receive_message_length': -1,
    'grpc.max_send_message_length': -1,
  });

  const cfg = await initializeConfig();
  const bundlr = new NodeBundlr(
    cfg.arweave_config.bundlr_node,
    cfg.arweave_config.wallet_chain,
    cfg.wallet_key,
  );
  const arweaveService = new ArweaveService(bundlr, cfg.arweave_config);
  const storageServer = new StorageServer(arweaveService);

  server.addService(StorageGrpcService, storageServer);
  addReflection(server, proto_descriptor_path);

  // Server startup
  const domain = '0.0.0.0:' + cfg.server.port;
  server.bindAsync(
    domain,
    ServerCredentials.createInsecure(),
    (err: Error | null, bindPort: number) => {
      if (err) {
        throw err;
      }
      logger.info(`gRPC:Server:${bindPort}`, new Date().toLocaleString());
      server.start();
    },
  );

  return server;
};
