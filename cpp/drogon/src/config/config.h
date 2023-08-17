#ifndef __CONFIG__CONFIG_H__
#define __CONFIG__CONFIG_H__

#include <string>

class DatabaseConfig {
public:
  DatabaseConfig(const std::string& url, int poolSize);
  DatabaseConfig(const DatabaseConfig& other);

  const std::string& getUrl() const;
  int getPoolSize() const;

  static DatabaseConfig load();
private:
  std::string url;
  int poolSize;
};

class ServerConfig {
public:
  ServerConfig(const std::string& host, int port);
  ServerConfig(const ServerConfig& other);

  const std::string& getHost() const;
  int getPort() const;
  int getNumThreads() const;

  static ServerConfig load();
private:
  std::string host;
  int port;
  int numThreads;
};

class Config {
public:
  Config(const ServerConfig& server, const DatabaseConfig& db);

  const ServerConfig& getServer() const;
  const DatabaseConfig& getDatabase() const;

  static Config load();

  ~Config();
private:
  Config();
  DatabaseConfig* db;
  ServerConfig* server;
};

#endif
