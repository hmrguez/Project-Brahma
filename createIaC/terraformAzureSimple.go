package createIaC

var terraformazuresimple = `
# Configure the Azure provider
provider "azurerm" {
  features {}
}

# Create a resource group
resource "azurerm_resource_group" "ecommerce_rg" {
  name     = "ecommerce-rg"
  location = "West US"
}

# Create a virtual network
resource "azurerm_virtual_network" "ecommerce_vnet" {
  name                = "ecommerce-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.ecommerce_rg.location
  resource_group_name = azurerm_resource_group.ecommerce_rg.name
}

# Create a subnet for the virtual network
resource "azurerm_subnet" "public_subnet" {
  name                 = "public-subnet"
  address_prefixes     = ["10.0.1.0/24"]
  virtual_network_name = azurerm_virtual_network.ecommerce_vnet.name
  resource_group_name  = azurerm_resource_group.ecommerce_rg.name
}

# Create a network security group for the virtual network
resource "azurerm_network_security_group" "webapp_nsg" {
  name                = "webapp-nsg"
  location            = azurerm_resource_group.ecommerce_rg.location
  resource_group_name = azurerm_resource_group.ecommerce_rg.name

  security_rule {
    name                       = "http"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "80"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

# Create a public IP address
resource "azurerm_public_ip" "webapp_public_ip" {
  name                = "webapp-public-ip"
  location            = azurerm_resource_group.ecommerce_rg.location
  resource_group_name = azurerm_resource_group.ecommerce_rg.name
  allocation_method   = "Static"
}

# Create a load balancer
resource "azurerm_lb" "webapp_lb" {
  name                = "webapp-lb"
  location            = azurerm_resource_group.ecommerce_rg.location
  resource_group_name = azurerm_resource_group.ecommerce_rg.name

  frontend_ip_configuration {
    name                 = "webapp-feip"
    public_ip_address_id = azurerm_public_ip.webapp_public_ip.id
  }

  backend_address_pool {
    name = "webapp-bepool"
  }

  probe {
    name                      = "webapp-probe"
    protocol                  = "Http"
    request_path              = "/"
    port                      = 80
    interval_seconds          = 15
    number_of_probes          = 3
    load_balancing_rule_ids   = [azurerm_lb_rule.webapp_lb_rule.id]
  }

  tags = {
    Environment = "production"
  }
}

# Create a load balancer rule
resource "azurerm_lb_rule" "webapp_lb_rule" {
  name                     = "webapp-lb-rule"
  protocol                 = "Tcp"
  frontend_port            = 80
  backend_port             = 80
  frontend_ip_configuration_id = azurerm_lb.webapp_lb.frontend_ip_configuration[0].id
  backend_address_pool_id  = azurerm_lb.webapp_lb.backend_address_pool[0].id
}

# Create a virtual machine scale set
resource "azurerm_linux_virtual_machine_scale_set" "webapp_scale_set" {
  name                = "webapp-scale-set"
  location            = azurerm_resource_group.ecommerce_rg.location
  resource_group_name = azurerm_resource_group.ecommerce_rg.name
  sku                 = "Standard_B1ls"
  instances           = 3

  storage_profile_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  os_disk {
    name              = "webapp-os-disk"
    caching           = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  network_interface {
    name    = "webapp-nic"
    primary = true

    ip_configuration {
      name                          = "webapp-ipconfig"
      subnet_id                     = azurerm_subnet.public_subnet.id
      load_balancer_backend_address_pool_ids = [azurerm_lb.webapp_lb.backend_address_pool[0].id]

      public_ip_address_id = azurerm_public_ip.webapp_public_ip.id
    }
  }

  os_profile {
    computer_name_prefix = "webapp"
    admin_username       = "adminuser"
    admin_password       = "your_admin_password_here"
  }

# Create an Azure SQL Database server
resource "azurerm_sql_server" "ecommerce_db_server" {
name                         = "ecommerce-db-server"
resource_group_name          = azurerm_resource_group.ecommerce_rg.name
location                     = azurerm_resource_group.ecommerce_rg.location
version                      = "12.0"
administrator_login          = "adminuser"
administrator_login_password = "your_admin_password_here"
}

# Create a firewall rule for the Azure SQL Database server
resource "azurerm_sql_firewall_rule" "db_firewall_rule" {
name                = "db-firewall-rule"
resource_group_name = azurerm_resource_group.ecommerce_rg.name
server_name         = azurerm_sql_server.ecommerce_db_server.name
start_ip_address    = azurerm_subnet.db_subnet.address_prefixes[0]
end_ip_address      = azurerm_subnet.db_subnet.address_prefixes[0]
}

# Create an Azure SQL Database
resource "azurerm_sql_database" "ecommerce_db" {
name                = "ecommerce-db"
resource_group_name = azurerm_resource_group.ecommerce_rg.name
server_name         = azurerm_sql_server.ecommerce_db_server.name
edition             = "Standard"
collation           = "SQL_Latin1_General_CP1_CI_AS"
max_size_gb         = 1
create_mode         = "Default"
}

# Grant the web subnet access to the Azure SQL Database
resource "azurerm_subnet_network_security_group_association" "web_subnet_nsg_association" {
subnet_id                 = azurerm_subnet.web_subnet.id
network_security_group_id = azurerm_network_security_group.web_nsg.id
}

# Grant the Azure SQL Database access to the web subnet
resource "azurerm_subnet_network_security_group_association" "db_subnet_nsg_association" {
subnet_id                 = azurerm_subnet.db_subnet.id
network_security_group_id = azurerm_sql_server.ecommerce_db_server.network_security_group_id
}

`