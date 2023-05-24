import 'source-map-support/register';
import 'reflect-metadata';
import { bootstrap } from '../config/bootstrap';
import { logger } from '../util';

function main() {
  bootstrap().catch((err) => {
    logger.error(`Server startup error, shut down: ${err}.`);
    process.kill(process.pid, 'SIGTERM');
  });
}

main();
