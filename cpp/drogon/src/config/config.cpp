#include "config.h"
#include <cstdlib>
#include <thread>

ServerConfig::ServerConfig(const std::string& host, int port)
  : host(host), port(port) {}

ServerConfig::ServerConfig(const ServerConfig& other)
  : host(other.host), port(other.port) {}

const std::string& ServerConfig::getHost() const {
  return host;
}

int ServerConfig::getPort() const {
  return port;
}

int ServerConfig::getNumThreads() const {
  return numThreads;
}

ServerConfig ServerConfig::load() {
  const char* host = std::getenv("SERVER__HOST");
  if (host == nullptr) {
    host = "localhost";
  }

  int port = 4000;
  const char* portStr = std::getenv("SERVER__PORT");
  if (portStr != nullptr) {
    int newPort = atoi(portStr);
    if (newPort > 0) {
      port = newPort;
    }
  }

  int numThreads = 1;
  const char* numThreadsStr = std::getenv("SERVER__NUM_THREADS");
  if (numThreadsStr != nullptr) {
    int newNumThreads = atoi(numThreadsStr);
    if (newNumThreads > 0) {
      numThreads = newNumThreads;
    }
  } else {
    numThreads = std::thread::hardware_concurrency() - 1;
    if (numThreads < 1) {
      numThreads = 1;
    }
  }

  return ServerConfig(host, port);
}


DatabaseConfig::DatabaseConfig(const std::string& url, int poolSize)
  : url(url), poolSize(poolSize) {}

DatabaseConfig::DatabaseConfig(const DatabaseConfig& other)
  : url(other.url) {}

const std::string& DatabaseConfig::getUrl() const {
  return url;
}

int DatabaseConfig::getPoolSize() const {
  return poolSize;
}

DatabaseConfig DatabaseConfig::load() {
  const char* url = std::getenv("DATABASE__URL");
  if (url == nullptr) {
    url = "postgres://restbench:restbench@localhost:5432/restbench";
  }

  int poolSize = 10;
  const char* poolSizeStr = std::getenv("DATABASE__POOL_SIZE");
  if (poolSizeStr != nullptr) {
    int newPoolSize = atoi(poolSizeStr);
    if (newPoolSize > 0) {
      poolSize = newPoolSize;
    }
  }

  return DatabaseConfig(url, poolSize);
}


Config::Config(const ServerConfig& server, const DatabaseConfig& db)
  : server(new ServerConfig(server)), db(new DatabaseConfig(db)) {}

Config::~Config() {
  delete this->server;
  delete this->db;
}

const ServerConfig& Config::getServer() const {
  return *this->server;
}

const DatabaseConfig& Config::getDatabase() const {
  return *this->db;
}

Config Config::load() {
  auto server = ServerConfig::load();
  auto database = DatabaseConfig::load();
  return Config(server, database);
}
