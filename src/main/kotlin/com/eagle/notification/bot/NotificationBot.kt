package com.eagle.notification.bot

import com.eagle.notification.bot.ui.NOT_AUTH
import com.eagle.notification.bot.ui.Text
import dev.inmo.tgbotapi.bot.TelegramBot
import dev.inmo.tgbotapi.bot.exceptions.CommonBotException
import dev.inmo.tgbotapi.extensions.behaviour_builder.buildBehaviour
import dev.inmo.tgbotapi.extensions.behaviour_builder.buildBehaviourWithLongPolling
import dev.inmo.tgbotapi.extensions.behaviour_builder.triggers_handling.onCommand
import io.ktor.client.plugins.*
import kotlinx.coroutines.CancellationException
import kotlinx.coroutines.CoroutineScope

class NotificationBot(
    private val bot: TelegramBot
) : AutoCloseable {


    suspend fun run(coroutineScope: CoroutineScope) {
        bot.buildBehaviourWithLongPolling(defaultExceptionsHandler = { error ->
            if (error !is CancellationException) {
                if (isUpdateTimeoutException(error)) {
//                    logger.warn(error) { "Timeout exception" }
                } else {
//                    logger.error(error) { "Unexpected error" }
                }
            }
        }) {
            onCommand("start") { msg ->
                val chatId = msg.chat.id
                //get chatId and check auth
                val notAuth = false
                if (notAuth) {
                    val text = Text(NOT_AUTH)
                    text.publish(chatId, bot)
                    return@onCommand
                }


            }
        }


    }

    override fun close() {
        TODO("Not yet implemented")
    }

    private fun isUpdateTimeoutException(e: Throwable): Boolean {
        val timeoutException = e as? HttpRequestTimeoutException
            ?: (e as? CommonBotException)?.cause as? HttpRequestTimeoutException

        return timeoutException != null && (timeoutException.message?.contains("getUpdates") == true)
    }
}