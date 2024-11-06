CREATE TRIGGER trigger_calcula_pontos
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION calcula_pontos();

CREATE TRIGGER trigger_atualiza_disponibilidade_prato
AFTER UPDATE ON ingrediente
FOR EACH ROW
EXECUTE FUNCTION atualiza_disponibilidade_prato();

CREATE TRIGGER trigger_impedir_compra_prato_indisponivel
BEFORE INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION impedir_compra_prato_indisponivel();

CREATE TRIGGER trigger_atualiza_quantidade_ingrediente
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION atualiza_quantidade_ingrediente();
