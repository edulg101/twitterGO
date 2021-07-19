$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("teste")



    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops.. ", "As senhas não coincidem", "error");
        return;
    }
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),

        }
    }).done(function () {
        Swal.fire("Sucesso ! ", "Usuário Cadastrado com Sucesso", "success")
            .then(() => {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(() => {
                    window.location = "/home";
                }).fail(() => {
                    Swal.fire("Ops...", "Erro ao autenticar usuario", "error");
                })
            });
    }).fail(function (erro) {
        Swal.fire("Ops..... ", "erro ao cadastrar usuario", "error");
    });
}