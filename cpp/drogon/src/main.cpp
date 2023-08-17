#include <iostream>
#include <drogon/drogon.h>

#include "config/config.h"
#include "transcript/transcript.h"

int main(int argc, char** argv) {
  const auto config = Config::load();

  const auto db = drogon::orm::DbClient::newPgClient(config.getDatabase().getUrl(), config.getDatabase().getPoolSize());

  const auto transcriptService = std::make_shared<transcript::TranscriptService>(db);
  const auto transcriptController = std::make_shared<api::v1::transcript::TranscriptController>(transcriptService);

  drogon::app()
    .setLogLevel(trantor::Logger::kDebug)
    .addListener(config.getServer().getHost(), config.getServer().getPort())
    .setThreadNum(config.getServer().getNumThreads())
    .run();

  return 0;
}
