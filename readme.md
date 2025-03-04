Auth Controller
Login: Para autenticar o usuário.
Logout: Para deslogar o usuário.
RefreshToken: Para renovar o token de autenticação.

User Controller
GetUserDetails: Para obter informações do perfil do usuário.
UpdateUserDetails: Para o usuário atualizar seu perfil.
ChangePassword: Para permitir que o usuário altere sua senha.
DeleteUser: Para deletar o usuário.

Product Controller
ListProducts: Para listar todos os produtos disponíveis na loja (exemplo: camisetas).
GetProductById: Para buscar um produto específico por ID.

Cart Controller
AddProductToCart: Para adicionar um produto ao carrinho.
GetCart: Para visualizar os itens no carrinho.
RemoveProductFromCart: Para remover um produto do carrinho.
ClearCart: Para limpar o carrinho de compras.

Order Controller
CreateOrder: Para criar um pedido com os itens do carrinho.
GetOrderDetails: Para obter detalhes de um pedido específico.
CancelOrder: Para cancelar um pedido (caso seja permitido).
ListOrders: Para listar todos os pedidos feitos pelo usuário.
TrackOrder: Para rastrear o status da entrega do pedido.

Payment Controller
CreatePaymentIntent: Para iniciar o processo de pagamento via Stripe Checkout (gerar o pagamento).
WebhookPaymentStatus: Para receber atualizações de status de pagamento do Stripe.

Shipping Controller
GetShippingOptions: Para obter opções de envio disponíveis (com integração com LOGGI).
CreateShipping: Para criar o pedido de envio com LOGGI.
