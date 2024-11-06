-- Triggers

-- Trigger 1: Calcula Pontos

CREATE TRIGGER trigger_calcula_pontos
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION calcula_pontos();

-- Trigger 2: Atualizar disponibilidade de pratos com ingredientes vencidos

CREATE TRIGGER trigger_atualiza_disponibilidade_prato
AFTER UPDATE ON ingrediente
FOR EACH ROW
EXECUTE FUNCTION atualiza_disponibilidade_prato();

-- Trigger 3: Impedir compra de prato indisponível

CREATE TRIGGER trigger_impedir_compra_prato_indisponivel
BEFORE INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION impedir_compra_prato_indisponivel();

-- Trigger 4: Atualizar a quantidade de ingredientes de venda

CREATE TRIGGER trigger_atualiza_quantidade_ingrediente
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION atualiza_quantidade_ingrediente();