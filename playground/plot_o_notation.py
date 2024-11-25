import matplotlib.pyplot as plt
import numpy as np

# Данные для графиков
n = np.arange(1, 101)
o1 = np.ones(100)
ologn = np.log2(n)
on = n
onlogn = n * np.log2(n)
on2 = n ** 2
o2n = 2 ** n

# Построение графиков
plt.figure(figsize=(10, 6))
plt.plot(n, o1, label='O(1)')
plt.plot(n, ologn, label='O(log n)')
plt.plot(n, on, label='O(n)')
plt.plot(n, onlogn, label='O(n log n)')
plt.plot(n, on2, label='O(n^2)')
plt.plot(np.arange(1, 11), o2n[:10], label='O(2^n)')  # Ограничиваем до 10, чтобы избежать слишком больших значений

plt.xlabel('n')
plt.ylabel('Time')
plt.title('O-нотация')
plt.legend()
plt.grid(True)
plt.yscale('log')  # Используем логарифмическую шкалу для оси y, чтобы лучше видеть различия
plt.show()
