CREATE DATABASE IF NOT EXISTS `petshop` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `petshop`;
-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: 192.168.66.102    Database: petshop
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Animal`
--

DROP TABLE IF EXISTS `Animal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Animal` (
  `idAnimal` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  `Porte_idPorte` int NOT NULL,
  `Raca_idRaca` int NOT NULL,
  `cliente_idcliente` int NOT NULL,
  PRIMARY KEY (`idAnimal`,`cliente_idcliente`),
  KEY `fk_Anima_Porte_idx` (`Porte_idPorte`),
  KEY `fk_Anima_Raca1_idx` (`Raca_idRaca`),
  KEY `fk_Anima_cliente1_idx` (`cliente_idcliente`),
  CONSTRAINT `fk_Anima_cliente1` FOREIGN KEY (`cliente_idcliente`) REFERENCES `cliente` (`idcliente`),
  CONSTRAINT `fk_Anima_Porte` FOREIGN KEY (`Porte_idPorte`) REFERENCES `Porte` (`idPorte`),
  CONSTRAINT `fk_Anima_Raca1` FOREIGN KEY (`Raca_idRaca`) REFERENCES `Raca` (`idRaca`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Animal`
--

LOCK TABLES `Animal` WRITE;
/*!40000 ALTER TABLE `Animal` DISABLE KEYS */;
/*!40000 ALTER TABLE `Animal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Cliente`
--

DROP TABLE IF EXISTS `Cliente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Cliente` (
  `idcliente` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  `cpf` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idcliente`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Cliente`
--

LOCK TABLES `Cliente` WRITE;
/*!40000 ALTER TABLE `Cliente` DISABLE KEYS */;
/*!40000 ALTER TABLE `Cliente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Funcionario`
--

DROP TABLE IF EXISTS `Funcionario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Funcionario` (
  `idFuncionario` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  `cpf` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idFuncionario`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Funcionario`
--

LOCK TABLES `Funcionario` WRITE;
/*!40000 ALTER TABLE `Funcionario` DISABLE KEYS */;
/*!40000 ALTER TABLE `Funcionario` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Pagamento`
--

DROP TABLE IF EXISTS `Pagamento`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Pagamento` (
  `idFormaPagamento` int NOT NULL AUTO_INCREMENT,
  `TipoPagamento_idTipoPagamento` int NOT NULL,
  `data` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`idFormaPagamento`),
  KEY `fk_Pagamento_TipoPagamento1_idx` (`TipoPagamento_idTipoPagamento`),
  CONSTRAINT `fk_Pagamento_TipoPagamento1` FOREIGN KEY (`TipoPagamento_idTipoPagamento`) REFERENCES `TipoPagamento` (`idTipoPagamento`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Pagamento`
--

LOCK TABLES `Pagamento` WRITE;
/*!40000 ALTER TABLE `Pagamento` DISABLE KEYS */;
/*!40000 ALTER TABLE `Pagamento` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Porte`
--

DROP TABLE IF EXISTS `Porte`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Porte` (
  `idPorte` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idPorte`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Porte`
--

LOCK TABLES `Porte` WRITE;
/*!40000 ALTER TABLE `Porte` DISABLE KEYS */;
INSERT INTO `Porte` VALUES (1,'Pequeno'),(2,'Médio'),(3,'Grande');
/*!40000 ALTER TABLE `Porte` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Produto`
--

DROP TABLE IF EXISTS `Produto`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Produto` (
  `idProduto` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) NOT NULL,
  `preço` float NOT NULL,
  PRIMARY KEY (`idProduto`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Produto`
--

LOCK TABLES `Produto` WRITE;
/*!40000 ALTER TABLE `Produto` DISABLE KEYS */;
INSERT INTO `Produto` VALUES (1,'Banho e tosa',100),(2,'Banhho',60),(3,'Bolinha',5),(4,'Cama  50cm x 100cm',70),(5,'Cama  50cm x 50cm',50),(6,'Cama  70cm x 70cm',60),(7,'Shampoo 500ml',18.32),(8,'Shampoo 1000ml',25.5),(9,'Consultas ',80.5);
/*!40000 ALTER TABLE `Produto` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Raca`
--

DROP TABLE IF EXISTS `Raca`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Raca` (
  `idRaca` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idRaca`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Raca`
--

LOCK TABLES `Raca` WRITE;
/*!40000 ALTER TABLE `Raca` DISABLE KEYS */;
INSERT INTO `Raca` VALUES (1,'Pastor-alemão'),(2,'Buldogue'),(3,'Labrador retriever'),(4,'Golden retriever'),(5,'Husky siberiano'),(6,'Shih-tzu');
/*!40000 ALTER TABLE `Raca` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Servico`
--

DROP TABLE IF EXISTS `Servico`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Servico` (
  `idServiço` int NOT NULL AUTO_INCREMENT,
  `Funcionario_idFuncionario` int NOT NULL,
  `data` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `TipoServico_idTipoServico` int NOT NULL,
  `Anima_idAnima` int DEFAULT NULL,
  `Anima_cliente_idcliente` int NOT NULL,
  `Produto_idProduto` int NOT NULL,
  `Pagamento_idFormaPagamento` int NOT NULL,
  PRIMARY KEY (`idServiço`,`Pagamento_idFormaPagamento`),
  KEY `fk_Serviço_Funcionario1_idx` (`Funcionario_idFuncionario`),
  KEY `fk_Serviço_TipoServico1_idx` (`TipoServico_idTipoServico`),
  KEY `fk_Serviço_Anima1_idx` (`Anima_idAnima`,`Anima_cliente_idcliente`),
  KEY `fk_Serviço_Produto1_idx` (`Produto_idProduto`),
  KEY `fk_Serviço_Pagamento1_idx` (`Pagamento_idFormaPagamento`),
  CONSTRAINT `fk_Serviço_Anima1` FOREIGN KEY (`Anima_idAnima`, `Anima_cliente_idcliente`) REFERENCES `Animal` (`idAnimal`, `cliente_idcliente`),
  CONSTRAINT `fk_Serviço_Funcionario1` FOREIGN KEY (`Funcionario_idFuncionario`) REFERENCES `Funcionario` (`idFuncionario`),
  CONSTRAINT `fk_Serviço_Pagamento1` FOREIGN KEY (`Pagamento_idFormaPagamento`) REFERENCES `Pagamento` (`idFormaPagamento`),
  CONSTRAINT `fk_Serviço_Produto1` FOREIGN KEY (`Produto_idProduto`) REFERENCES `Produto` (`idProduto`),
  CONSTRAINT `fk_Serviço_TipoServico1` FOREIGN KEY (`TipoServico_idTipoServico`) REFERENCES `TipoServico` (`idTipoServico`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Servico`
--

LOCK TABLES `Servico` WRITE;
/*!40000 ALTER TABLE `Servico` DISABLE KEYS */;
/*!40000 ALTER TABLE `Servico` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TipoPagamento`
--

DROP TABLE IF EXISTS `TipoPagamento`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TipoPagamento` (
  `idTipoPagamento` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idTipoPagamento`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TipoPagamento`
--

LOCK TABLES `TipoPagamento` WRITE;
/*!40000 ALTER TABLE `TipoPagamento` DISABLE KEYS */;
INSERT INTO `TipoPagamento` VALUES (1,'Dinheiro'),(2,'Cartão'),(3,'Boleto');
/*!40000 ALTER TABLE `TipoPagamento` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TipoServico`
--

DROP TABLE IF EXISTS `TipoServico`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TipoServico` (
  `idTipoServico` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idTipoServico`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TipoServico`
--

LOCK TABLES `TipoServico` WRITE;
/*!40000 ALTER TABLE `TipoServico` DISABLE KEYS */;
INSERT INTO `TipoServico` VALUES (1,'Venda'),(2,'Compra');
/*!40000 ALTER TABLE `TipoServico` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
